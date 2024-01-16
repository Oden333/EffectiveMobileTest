package repository

import (
	"EMtest/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User_repo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *User_repo {
	return &User_repo{db: db}
}

// TODO: +age +sex +country
func (r *User_repo) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("Insert into %s (name,password) values ($1, $2) returning id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
