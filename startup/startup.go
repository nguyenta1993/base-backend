package startup

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"base_service/config"
	"base_service/database"
	"base_service/internal"
	"base_service/internal/api"
	v "base_service/internal/validation"
	"base_service/pkg/healthcheck"
	"base_service/pkg/metrics"

	// Just exmple, in the real world, using https://github.com/gogovan-korea/s14e-backend-proto
	proto "base_service/internal/api/grpc/proto_gen"

	"github.com/gammazero/workerpool"
	r "github.com/go-redis/redis/v8"
	"github.com/gogovan-korea/ggx-kr-service-utils/command"
	"github.com/gogovan-korea/ggx-kr-service-utils/grpc/client"
	"github.com/gogovan-korea/ggx-kr-service-utils/kafka"
	"github.com/gogovan-korea/ggx-kr-service-utils/localization"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/gogovan-korea/ggx-kr-service-utils/redis"
	"github.com/gogovan-korea/ggx-kr-service-utils/tracing"
	"github.com/gogovan-korea/ggx-kr-service-utils/validation"
	k "github.com/segmentio/kafka-go"
)

func runServer(
	ctx context.Context,
	logger logger.Logger,
	container *api.ApiContainer,
	readDb database.ReadDb,
	writeDb database.WriteDb,
	redisClient r.UniversalClient,
	kafkaConn *k.Conn,
) {
	wp := workerpool.New(4)
	// Run healthcheck
	wp.Submit(healthcheck.RunHealthCheck(ctx, logger, cfg, readDb, writeDb, redisClient, kafkaConn))
	// Run metrics
	wp.Submit(metrics.RunMetrics(logger, cfg))
	// Run Grpc server
	wp.Submit(container.GrpcServer.Run)
	// Run Http server
	wp.Submit(container.HttpServer.Run)

	wp.StopWait()
}

func registerDependencies(ctx context.Context, logger logger.Logger) (*api.ApiContainer, database.ReadDb, database.WriteDb, r.UniversalClient) {
	// Open database connection
	readDb, writeDb := database.Open(cfg.Database, logger)
	// Register kafka
	kafkaProducer := kafka.NewProducer(logger, &k.Writer{
		Addr:         k.TCP(cfg.Kafka.Config.Brokers...),
		BatchTimeout: 1 * time.Nanosecond, // default 1 second
		Async:        true,                // default false
	})
	// Register user service client conn
	// It's just example, it calls itself
	userServiceClient := proto.NewUserServiceClient(client.NewClientConn(ctx, logger, cfg.GRPC.Port, true))
	// Register redis client
	redisClient := redis.NewUniversalRedisClient(redis.Config(*cfg.Redis))
	// Register dependencies
	return internal.InitializeContainer(cfg, logger, kafkaProducer, userServiceClient, redisClient, &readDb, &writeDb),
		readDb,
		writeDb,
		redisClient
}

var cfg *config.AppConfig

func Execute() {
	// Init AppConfig
	cfg = &config.AppConfig{}

	// Init commands
	command.UseCommands(
		command.WithStartCommand(start, cfg, "database.writeDb"),
	)
}

func start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	// Init logger
	logger := logger.NewZapLogger(cfg.Logger.LogLevel)
	// Register dependencies
	container, readDb, writeDb, redisClient := registerDependencies(ctx, logger)
	// Init resources for localization
	localization.InitResources(cfg.Http.Resources)
	// Init kakfa
	kafkaConn := kafka.UseKafka(
		ctx,
		logger,
		&kafka.Config{
			Config: (*kafka.ConfigDetail)(cfg.Kafka.Config),
			Topics: []kafka.TopicConfig{
				kafka.TopicConfig(cfg.Kafka.Topics.UserUpdated),
			},
		},
		&kafka.ConsumerConfig{
			Topics:   []string{cfg.Kafka.Topics.UserUpdated.TopicName},
			PoolSize: cfg.Kafka.Config.NumWorker,
			Worker:   container.Consumer.ProcessMessages,
		},
	)
	// Init tracing
	tracing.UseOpenTelemetry(tracing.Config(*cfg.Jaeger))
	// Init validation
	validation.UseValidation(validation.CustomValidation{Tag: v.Tag, ValidatorFunc: v.AgeNotNegative})
	// Run server
	runServer(ctx, logger, container, readDb, writeDb, redisClient, kafkaConn)
}
