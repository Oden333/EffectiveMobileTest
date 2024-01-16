package service

import (
	"EMtest/models"
	"EMtest/pkg/repository"
)

type User_service struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *User_service {
	return &User_service{repo: repo}
}

func (s *User_service) CreateUser(user models.User) (int, error) {
	//TODO: апи доп инфа
	return s.repo.CreateUser(user)
}
