package repository

import (
	taskmanagement "taskmanager"

	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) createTask(task taskmanagement.Tasks) (int, error) {
	return 0, nil
}
