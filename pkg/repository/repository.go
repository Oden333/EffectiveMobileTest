package repository

import (
	"EMtest/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(user models.User) (int, error)
}

type Repository struct {
	UserRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(db),
	}
}
