package repository

import (
	"EMtest/models"
	helpers "EMtest/pkg/handler/helper"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(user models.User) (int, error)
	GetAllUsers(limit int, offset int) (int, []models.User, error)
	GetCertainUsers(limit int, offset int, filter helpers.FilterData) (int, []models.User, error)
	DeleteUser(userId int) error
}

type Repository struct {
	UserRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db),
	}
}
