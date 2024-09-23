package main

import (
	"log"
	"os"
	taskmanagement "taskmanager"
	"taskmanager/pkg/handler"
	"taskmanager/pkg/repository"
	"taskmanager/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variabler: %s", err.Error())
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error connecting to database")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(taskmanagement.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server : %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
