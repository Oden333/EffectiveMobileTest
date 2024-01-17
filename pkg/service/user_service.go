package service

import (
	"EMtest/models"
	"EMtest/pkg/repository"
	"log"
)

type User_service struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *User_service {
	return &User_service{repo: repo}
}

func (s *User_service) CreateUser(user models.User) (int, error) {
	//TODO: апи доп инфа

	//Запрашиваем вероятную национальность
	country, err := GetCountry(user.Name)
	if err != nil && err != ErrorEmptyCountry {
		log.Printf("Error occured while getting country info, %s", err)
		return -1, err
	}
	if err == ErrorEmptyCountry {
		country = "No country info"
	}
	user.Country = country

	//Запрашиваем возраст
	age, err := GetAge(user.Name)
	if err != nil {
		log.Printf("Error occured while getting age info, %s", err)
		return -1, err
	}
	user.Age = age

	//Запрашиваем пол
	sex, err := GetGender(user.Name)
	if err != nil {
		log.Printf("Error occured while getting gender info, %s", err)
		return -1, err
	}
	user.Gender = sex

	return s.repo.CreateUser(user)
}
