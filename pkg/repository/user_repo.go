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
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	query := fmt.Sprintf("Insert into %s (name, surname, patronymic, age, gender, country) VALUES ($1,$2,$3,$4,$5,$6) returning id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Country)
	if err != row.Scan(&id) {
		tx.Rollback()
		return 0, err
	}
	return id, nil
}
