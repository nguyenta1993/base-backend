//Interface that will be implemented in infrastructure
package user

import (
	"base_service/internal/domain/entities"
	"context"
)

type UserCommandRepository interface {
	UpdateUser(ctx context.Context, user *entities.User) (bool, error)
	CreateUser(ctx context.Context, user *entities.User) (bool, error)
}
