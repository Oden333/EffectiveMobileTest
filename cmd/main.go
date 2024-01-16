package main

import (
	srv "EMtest"
	"EMtest/pkg/handler"
	"EMtest/pkg/repository"
	"EMtest/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	//Инит логгер
	logrus.SetFormatter(new(logrus.JSONFormatter))
	//Инит конфиг данных через енв переменные
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("unable to get pswd (%s)", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db (%s)", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(srv.Server)
	go func() {

		if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("App started")
}
