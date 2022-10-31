package createuser

import (
	"base_service/config"
	"base_service/internal/domain/entities"
	interfaces "base_service/internal/domain/interfaces/user"
	"context"

	k "github.com/gogovan-korea/ggx-kr-service-utils/kafka"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type CreateUserHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	repo          interfaces.UserCommandRepository
	redisRepo     interfaces.CacheRepository
	kafkaProducer *k.Producer
}

func NewCreateUserHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.UserCommandRepository,
	redisRepo interfaces.CacheRepository,
	kafkaProducer *k.Producer,
) *CreateUserHandler {
	return &CreateUserHandler{logger, cfg, repo, redisRepo, kafkaProducer}
}

func (h *CreateUserHandler) Handle(ctx context.Context, createUserCommand *CreateUserCommand) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateUser.Handle")
	defer span.Finish()

	user := entities.User(*createUserCommand)
	_, err := h.repo.CreateUser(ctx, &user)

	if err != nil {
		h.logger.Error("User has not been created successfully", zap.Error(err))
		return false, err
	}

	h.redisRepo.SetUser(ctx, &user, user.Username)
	return true, nil
}
