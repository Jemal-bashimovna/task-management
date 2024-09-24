package repository

import (
	"fmt"
	"log"
	taskmanagement "taskmanager"

	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) CreateTask(task taskmanagement.Tasks) (int, error) {

	var id int

	query := fmt.Sprintf("insert into %s (title, description) values ($1, $2) RETURNING id", tasksTable)

	row := r.db.QueryRow(query, task.Title, task.Description)
	fmt.Println(query)
	if err := row.Scan(&id); err != nil {
		log.Printf("failed to scan id : %v", err)
		return 0, err
	}

	return id, nil
}
