package updateuser

import (
	"context"
	"time"

	"base_service/config"
	usermessages "base_service/internal/api/kafka/proto_gen"
	"base_service/internal/domain/entities"
	interfaces "base_service/internal/domain/interfaces/user"

	"github.com/gammazero/workerpool"
	k "github.com/gogovan-korea/ggx-kr-service-utils/kafka"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/gogovan-korea/ggx-kr-service-utils/tracing"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type UpdateUserHandler struct {
	logger        logger.Logger
	cfg           *config.AppConfig
	repo          interfaces.UserCommandRepository
	redisRepo     interfaces.CacheRepository
	kafkaProducer *k.Producer
}

func NewUpdateUserHandler(
	logger logger.Logger,
	cfg *config.AppConfig,
	repo interfaces.UserCommandRepository,
	redisRepo interfaces.CacheRepository,
	kafkaProducer *k.Producer,
) *UpdateUserHandler {
	return &UpdateUserHandler{logger, cfg, repo, redisRepo, kafkaProducer}
}

func (h *UpdateUserHandler) Handle(ctx context.Context, updateUserCommand *UpdateUserCommand) (bool, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "UpdateUser.Handle")
	defer span.End()

	// Update database
	user := entities.User(*updateUserCommand)
	_, err := h.repo.UpdateUser(ctx, &user)

	if err != nil {
		h.logger.Error("User has not been updated successfully", zap.Error(err))
		return false, err
	}

	wp := workerpool.New(2)

	// Update cache
	wp.Submit(func() {
		if err := h.redisRepo.SetUser(ctx, &user, user.Username); err != nil {
			h.logger.Error("redisRepo.SetUser", zap.Error(err))
		}
	})

	// Kafka
	wp.Submit(func() {
		userMessage := &usermessages.UserUpdated{
			Username:    user.Username,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}

		userMessageBytes, _ := proto.Marshal(userMessage)
		message := kafka.Message{
			Topic:   h.cfg.Kafka.Topics.UserUpdated.TopicName,
			Value:   userMessageBytes,
			Time:    time.Now().UTC(),
			Headers: tracing.GetKafkaTracingHeadersFromCtx(ctx),
		}

		h.logger.Info("Start to send message to Kafka: " + time.Now().String())
		if err = h.kafkaProducer.PublishMessage(ctx, message); err != nil {
			h.logger.Error("kafkaProducer.PublishMessage", zap.Error(err))
		}
	})

	wp.StopWait()

	return true, nil
}
