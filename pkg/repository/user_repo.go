package repository

import (
	"EMtest/models"
	helpers "EMtest/pkg/handler/helper"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type User_repo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *User_repo {
	return &User_repo{db: db}
}

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
func (r *User_repo) GetAllUsers(limit int, offset int) (int, []models.User, error) {

	//Узнаём общее кол-во человек (для настроек странциы)
	var count int
	err := r.db.Get(&count, fmt.Sprintf("SELECT COUNT(*) FROM %s", usersTable))
	if err != nil {
		logrus.Debug("Error while making Users-counting DB request ")
		return -1, nil, err
	}

	//Выбираем людей
	var users []models.User
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"
	err = r.db.Select(&users, query, limit, offset)
	if err != nil {
		return -1, nil, err
	}

	logrus.Info("Got response from DB for GetAll request")

	return count, users, nil
}

func (r *User_repo) GetCertainUsers(limit int, offset int, filter helpers.FilterData) (int, []models.User, error) {

	// Базовый запрос
	// 1=1 для более удобного формирования конструкции с фильтром
	query := `
        SELECT * FROM users
        WHERE 1=1
    `
	//Узнаём, какие фильтры будут добавляться и формируем запрос в бд
	if filter.Age != "" {
		query += " AND age = " + filter.Age
	}
	if filter.Name != "" {
		query += " AND name = '" + filter.Name + "'"
	}
	if filter.Surname != "" {
		query += " AND surname = '" + filter.Surname + "'"
	}
	if filter.Patronymic != "" {
		query += " AND patronymic = '" + filter.Patronymic + "'"
	}
	if filter.Country != "" {
		query += " AND country = '" + filter.Country + "'"
	}

	// Так же добавим лимит и оффсет для пагинации страниц
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	/*
			var count int
		   	//подсчитывает количество записей, удовлетворяющих условиям в основном запросе.
		   	if err := r.db.Get(&count, "SELECT COUNT (*) FROM ("+query+")"); err != nil {
		   		return 0, nil, err
		   	}
	*/

	var users []models.User
	if err := r.db.Select(&users, query); err != nil {
		return 0, nil, err
	}

	count := len(users)

	return count, users, nil
}

func (r *User_repo) DeleteUser(userId int) error {

	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		logrus.Debug("Error deleting User by Id", err)

		return err
	}
	return nil
}
