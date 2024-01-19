package repository

import (
	"EMtest/models"
	helpers "EMtest/pkg/handler/helper"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(user models.User) (int, error)
	GetAllUsers(limit int, offset int) (int, []models.User, error)
	GetCertainUsers(limit int, offset int, filter map[string]string) (int, []models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user helpers.UserData) error
}

type Repository struct {
	UserRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db),
	}
}
