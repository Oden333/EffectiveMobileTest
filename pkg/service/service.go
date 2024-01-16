package service

import (
	"EMtest/models"
	"EMtest/pkg/repository"
)

type UserService interface {
	CreateUser(user models.User) (int, error)
}

type Service struct {
	UserService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repos.UserRepo),
	}
}
