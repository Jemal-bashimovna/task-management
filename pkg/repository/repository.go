package repository

import "github.com/jmoiron/sqlx"

type Tasks interface{}

type Repository struct {
	Tasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
