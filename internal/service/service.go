package service

import (
	userservice "base_service/internal/application/user"
)

type Service struct {
	UserService *userservice.UserService
}

func NewService(userService *userservice.UserService) *Service {
	return &Service{UserService: userService}
}
