package service

import (
	"EMtest/models"
	helpers "EMtest/pkg/handler/helper"
	"EMtest/pkg/repository"
)

type UserService interface {
	CreateUser(user models.User) (int, error)
	GetAllUsers(limit int, offset int) (int, []models.User, error)
	GetCertainUsers(limit int, offset int, filter helpers.FilterData) (int, []models.User, error)
	DeleteUser(userId int) error
}

type Service struct {
	UserService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repos.UserRepo),
	}
}
