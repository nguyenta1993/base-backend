package userservice

import (
	createuser "base_service/internal/application/user/commands/create_user"
	updateuser "base_service/internal/application/user/commands/update_user"
	getuser "base_service/internal/application/user/queries/get_user"
)

type UserService struct {
	// Commands
	UpdateUserHandler *updateuser.UpdateUserHandler
	CreateUserHandler *createuser.CreateUserHandler

	//Queries
	GetUserHandler *getuser.GetUserHandler
}

func NewUserService(
	getUserHandler *getuser.GetUserHandler,
	updateUserHandler *updateuser.UpdateUserHandler,
	createUserHandler *createuser.CreateUserHandler,
) *UserService {
	return &UserService{GetUserHandler: getUserHandler, UpdateUserHandler: updateUserHandler, CreateUserHandler: createUserHandler}
}
