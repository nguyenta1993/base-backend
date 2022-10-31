//Interface that will be implemented in infrastructure
package user

import (
	"base_service/internal/domain/entities"
	"context"
)

type UserQueryRepository interface {
	GetUser(ctx context.Context, username string) (*entities.User, error)
}
