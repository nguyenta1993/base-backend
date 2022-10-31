package kafka

import (
	"context"
	"time"

	usermessages "base_service/internal/api/kafka/proto_gen"
	updateuser "base_service/internal/application/user/commands/update_user"

	"github.com/avast/retry-go"
	"github.com/gogovan-korea/ggx-kr-service-utils/tracing"
	v "github.com/gogovan-korea/ggx-kr-service-utils/validation"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (c *Consumer) processUserUpdated(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "user_consumer.processUserUpdated")
	defer span.End()

	msg := &usermessages.UserUpdated{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		c.logger.Warn("proto.Unmarshal", zap.Error(err))
		c.commitErrMessage(ctx, r, m)
		return
	}

	updateUserCommand := &updateuser.UpdateUserCommand{Username: msg.GetUsername(), PhoneNumber: msg.GetPhoneNumber(), Email: msg.GetEmail()}
	if err := v.ValidateStruct(updateUserCommand); err != nil {
		c.logger.Error("Validate", zap.Error(err))
		c.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		// Add logic here
		//_, err := s.service.UserService.UpdateUserHandler.Handle(ctx, updateUserCommand)
		c.logger.Info("Consume kafka: " + time.Now().String())
		return nil
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		c.logger.Error("UpdateUserHandler.Handle", zap.Error(err))
		return
	}

	c.commitMessage(ctx, r, m)
}
