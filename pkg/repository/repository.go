package repository

import (
	taskmanagement "taskmanager"

	"github.com/jmoiron/sqlx"
)

type Tasks interface {
	CreateTask(task taskmanagement.Tasks) (int, error)
}

type Repository struct {
	Tasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Tasks: NewTaskPostgres(db),
	}
}
