package kafka

import (
	"context"
	"sync"

	"base_service/config"
	"base_service/internal/service"

	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Consumer struct {
	service *service.Service
	logger  logger.Logger
	cfg     *config.AppConfig
}

func NewConsumer(logger logger.Logger, cfg *config.AppConfig, service *service.Service) *Consumer {
	return &Consumer{logger: logger, cfg: cfg, service: service}
}

func (c *Consumer) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			c.logger.Warn("workerID", zap.Error(err))
			continue
		}

		switch m.Topic {
		case c.cfg.Kafka.Topics.UserUpdated.TopicName:
			c.processUserUpdated(ctx, r, m)
		}
	}
}

func (c *Consumer) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	c.logger.Info("Committed kafka message",
		zap.String("Topic", m.Topic),
		zap.Int("Partition", m.Partition),
		zap.Int64("Offset", m.Offset),
	)
	if err := r.CommitMessages(ctx, m); err != nil {
		c.logger.Error("commitMessage", zap.Error(err))
	}
}

func (c *Consumer) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	c.logger.Info("Committed kafka message",
		zap.String("Topic", m.Topic),
		zap.Int("Partition", m.Partition),
		zap.Int64("Offset", m.Offset),
	)
	if err := r.CommitMessages(ctx, m); err != nil {
		c.logger.Error("commitMessage", zap.Error(err))
	}
}
