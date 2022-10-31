//go:build wireinject
// +build wireinject

package internal

import (
	"base_service/config"
	database "base_service/database"
	"base_service/internal/api"
	grpcserver "base_service/internal/api/grpc"
	"base_service/internal/api/http"
	v1 "base_service/internal/api/http/v1"
	kafkaconsumer "base_service/internal/api/kafka"
	userservice "base_service/internal/application/user"
	createuser "base_service/internal/application/user/commands/create_user"
	updateuser "base_service/internal/application/user/commands/update_user"
	getuser "base_service/internal/application/user/queries/get_user"
	userrepo "base_service/internal/infrastructure/persistent/user"
	grpcmetrics "base_service/internal/metrics/grpc"
	httpmetrics "base_service/internal/metrics/http"
	service "base_service/internal/service"

	// Just exmple, in the real world, using https://github.com/gogovan-korea/s14e-backend-proto
	proto "base_service/internal/api/grpc/proto_gen"

	"github.com/go-redis/redis/v8"
	kafka "github.com/gogovan-korea/ggx-kr-service-utils/kafka"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/google/wire"
)

var container = wire.NewSet(
	api.NewApiContainer,
)

var apiSet = wire.NewSet(
	grpcserver.NewServer,
	http.NewServer,
	kafkaconsumer.NewConsumer,
)

var metricsSet = wire.NewSet(
	httpmetrics.NewHttpMetrics,
	grpcmetrics.NewGrpcMetrics,
)

var serviceSet = wire.NewSet(
	service.NewService,
	v1.NewUserHandler,
)

var specificServiceSet = wire.NewSet(
	userservice.NewUserService,
)

var handlerSet = wire.NewSet(
	getuser.NewGetUserHandler,
	updateuser.NewUpdateUserHandler,
	createuser.NewCreateUserHandler,
)

var repoSet = wire.NewSet(
	userrepo.NewUserQueryRepository,
	userrepo.NewUserCommandRepository,
	userrepo.NewUserRedisRepository,
)

func InitializeContainer(
	appCfg *config.AppConfig,
	logger logger.Logger,
	kafkaProducer *kafka.Producer,
	userServiceClient proto.UserServiceClient,
	redisClient redis.UniversalClient,
	readDb *database.ReadDb,
	writeDb *database.WriteDb,
) *api.ApiContainer {
	wire.Build(handlerSet, repoSet, specificServiceSet, serviceSet, metricsSet, apiSet, container)
	return &api.ApiContainer{}
}
