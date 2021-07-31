package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "go-boilerplate/docs"

	"github.com/spf13/viper"
	"go-boilerplate/internal/db"
	"go-boilerplate/internal/handlers"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/server"
	"go-boilerplate/internal/services"
	"log"
	"os"
)

// @title Go Boilerplate
// @version 1.0
// @description This is default RESTfulAPI boilerplate
// @host localhost:8000
// @basePath /
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	postgresDB, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repositories.NewRepository(postgresDB)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
