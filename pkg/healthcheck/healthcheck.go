package healthcheck

import (
	"context"
	"net/http"
	"time"

	"base_service/config"
	"base_service/database"
	"base_service/pkg/constants"

	"github.com/go-redis/redis/v8"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/heptiolabs/healthcheck"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func RunHealthCheck(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.AppConfig,
	readDb database.ReadDb,
	writeDb database.WriteDb,
	redisClient redis.UniversalClient,
	kafkaConn *kafka.Conn,
) func() {
	return func() {
		health := healthcheck.NewHandler()
		interval := time.Duration(cfg.Heathcheck.Interval) * time.Second

		livenessCheck(ctx, cfg, health)
		readinessCheck(ctx, logger, health, interval, readDb, writeDb, redisClient, kafkaConn)

		logger.Info("Heathcheck server listening on port", zap.String("Port", cfg.Heathcheck.Port))
		if err := http.ListenAndServe(cfg.Heathcheck.Port, health); err != nil {
			logger.Warn("Heathcheck server", zap.Error(err))
		}
	}
}

func livenessCheck(ctx context.Context, cfg *config.AppConfig, health healthcheck.Handler) {
	health.AddLivenessCheck(constants.GoroutineThreshold, healthcheck.GoroutineCountCheck(cfg.Heathcheck.GoroutineThreshold))
}

func readinessCheck(
	ctx context.Context,
	logger logger.Logger,
	health healthcheck.Handler,
	interval time.Duration,
	readDb database.ReadDb,
	writeDb database.WriteDb,
	redisClient redis.UniversalClient,
	kafkaConn *kafka.Conn,
) {

	health.AddReadinessCheck(constants.Redis, healthcheck.AsyncWithContext(ctx, func() error {
		err := redisClient.Ping(ctx).Err()
		return err
	}, interval))

	health.AddReadinessCheck(constants.ReadDatabase, healthcheck.AsyncWithContext(ctx, func() error {
		return readDb.DB.PingContext(ctx)
	}, interval))

	health.AddReadinessCheck(constants.WriteDatabase, healthcheck.AsyncWithContext(ctx, func() error {
		return writeDb.DB.PingContext(ctx)
	}, interval))
	health.AddReadinessCheck(constants.Kafka, healthcheck.AsyncWithContext(ctx, func() error {
		if _, err := kafkaConn.Brokers(); err != nil {
			return err
		}
		return nil
	}, interval))

}
