package user

import (
	"base_service/internal/domain/entities"
	"context"
)

type CacheRepository interface {
	SetUser(ctx context.Context, user *entities.User, key string) error
	GetUser(ctx context.Context, key string) *entities.User
}
