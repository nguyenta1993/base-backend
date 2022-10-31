// Implement user_repository for concrete databases
package userpersitent

import (
	"base_service/database"
	entities "base_service/internal/domain/entities"
	interfaces "base_service/internal/domain/interfaces/user"
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type userQueryRepository struct {
	*sqlx.DB
}

func NewUserQueryRepository(readDb *database.ReadDb) interfaces.UserQueryRepository {
	return &userQueryRepository{*readDb}
}

func (repo *userQueryRepository) GetUser(ctx context.Context, username string) (*entities.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userQueryRepository.GetUser")
	defer span.Finish()

	user := entities.User{}
	err := repo.GetContext(ctx, &user, "SELECT * FROM user WHERE username=?", username)

	return &user, err
}
