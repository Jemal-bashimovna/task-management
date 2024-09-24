package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	tasksTable = "tasks"
)

type Config struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func NewPostgres(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("can't connect to db")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("can't ping to db")
		return nil, err
	}

	return db, nil
}
