package updateuser

import (
	"context"
	"time"

	"base_service/config"
	usermessages "base_service/internal/api/kafka/proto_gen"
	interfaces "base_service/internal/domain/interfaces/user"

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

func (h *UpdateUserHandler) Handle(ctx context.Context) (bool, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "UpdateUser.Handle")
	defer span.End()

	// Update database

	userMessage := &usermessages.UserUpdated{
		Username:    "abc",
		Email:       "def",
		PhoneNumber: "123",
	}

	userMessageBytes, _ := proto.Marshal(userMessage)
	message := kafka.Message{
		Topic:   h.cfg.Kafka.Topics.UserUpdated.TopicName,
		Value:   userMessageBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromCtx(ctx),
	}

	h.logger.Info("Start to send message to Kafka: " + time.Now().String())
	if err := h.kafkaProducer.PublishMessage(ctx, message); err != nil {
		h.logger.Error("kafkaProducer.PublishMessage", zap.Error(err))
	}
	return true, nil
}
