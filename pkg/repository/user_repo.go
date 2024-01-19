package repository

import (
	"EMtest/models"
	helpers "EMtest/pkg/handler/helper"
	"fmt"
	"strings"

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

func (r *User_repo) UpdateUser(userId int, user helpers.UserData) error {
	//TODO: Добавить проверку существования ID в БД

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *&user.Name)
		argId++
	}

	if user.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *&user.Surname)
		argId++
	}

	if user.Patronymic != "" {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *&user.Patronymic)
		argId++
	}

	if user.Age != "" {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *&user.Age)
		argId++
	}

	if user.Gender != "" {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *&user.Gender)
		argId++
	}

	if user.Country != "" {
		setValues = append(setValues, fmt.Sprintf("country=$%d", argId))
		args = append(args, *&user.Country)
		argId++
	}

	// name=$1
	// в нужном порядке добавляем аргументы
	// surname=$_
	setQuery := strings.Join(setValues, ", ")

	//Формируем безопастный SQL запрос в бд
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, usersTable, setQuery, argId)

	args = append(args, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	//Выполняем запрос
	_, err := r.db.Exec(query, args...)

	return err

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

func (r *User_repo) GetCertainUsers(limit int, offset int, filter map[string]string) (int, []models.User, error) {

	// Базовый запрос
	// 1=1 для более удобного формирования конструкции с фильтром
	query := `
        SELECT * FROM users
        WHERE 1=1
    `
	//Узнаём, какие фильтры будут добавляться и формируем запрос в бд
	for key, value := range filter {
		query += fmt.Sprintf(" AND %s = '%s'", key, value)
	}

	fmt.Println(query)

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
	//В бд ID удаляется, но список остаётся без смещения
	// ID	Name ...
	// 11	ivan ...
	// 13	john ...
	if err != nil {
		logrus.Debug("Error deleting User by Id", err)

		return err
	}
	return nil
}
