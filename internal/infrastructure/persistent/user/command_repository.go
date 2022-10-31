package userpersitent

import (
	"base_service/database"
	entities "base_service/internal/domain/entities"
	interfaces "base_service/internal/domain/interfaces/user"
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type userCommandRepository struct {
	*sqlx.DB
	logger logger.Logger
}

func NewUserCommandRepository(writeDb *database.WriteDb, logger logger.Logger) interfaces.UserCommandRepository {
	return &userCommandRepository{*writeDb, logger}
}

func (repo *userCommandRepository) CreateUser(ctx context.Context, user *entities.User) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userCommandRepository.CreateUser")
	defer span.Finish()

	_, err := repo.NamedExec(`INSERT INTO user(username, email, phonenumber) VALUES (:username, :email, :phonenumber)`,
		map[string]interface{}{
			"username":    user.Username,
			"email":       user.Email,
			"phonenumber": user.PhoneNumber,
		},
	)

	if err != nil {
		repo.logger.Error("User has not been created successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}

func (repo *userCommandRepository) UpdateUser(ctx context.Context, user *entities.User) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userCommandRepository.UpdateUser")
	defer span.Finish()

	_, err := repo.NamedExec(`UPDATE user SET phonenumber=:phonenumber WHERE username=:username`,
		map[string]interface{}{
			"phonenumber": user.PhoneNumber,
			"username":    user.Username,
		},
	)

	if err != nil {
		repo.logger.Error("User has not been updated successfully", zap.Error(err))
		return false, err
	}

	return true, nil
}
