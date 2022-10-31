package userpersitent

import (
	"context"
	"encoding/json"

	entities "base_service/internal/domain/entities"
	interfaces "base_service/internal/domain/interfaces/user"

	"github.com/go-redis/redis/v8"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type userRedisRepository struct {
	logger      logger.Logger
	redisClient redis.UniversalClient
}

func NewUserRedisRepository(logger logger.Logger, redisClient redis.UniversalClient) interfaces.CacheRepository {
	return &userRedisRepository{logger, redisClient}
}

func (r *userRedisRepository) SetUser(ctx context.Context, user *entities.User, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRedisRepository.CreateUser")
	defer span.Finish()

	userBytes, _ := json.Marshal(user)
	if err := r.redisClient.Set(ctx, key, userBytes, 0).Err(); err != nil {
		r.logger.Error("userRedisRepository.Set", zap.Error(err))
	}

	return nil
}

func (r *userRedisRepository) GetUser(ctx context.Context, key string) *entities.User {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRedisRepository.GetUser")
	defer span.Finish()

	user, _ := r.redisClient.Get(ctx, key).Bytes()

	if user == nil {
		return nil
	}

	var u entities.User
	if err := json.Unmarshal(user, &u); err != nil {
		r.logger.Error("User.Unmarshal", zap.Error(err))
		return nil
	}

	return &u
}
