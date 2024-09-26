package repository

import (
	"fmt"
	"log"
	"strings"
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

	if err := row.Scan(&id); err != nil {
		log.Printf("failed to scan id : %v", err)
		return 0, err
	}

	return id, nil
}

func (r *TaskPostgres) GetTasks() ([]taskmanagement.Tasks, error) {
	
	var tasks []taskmanagement.Tasks

	query := fmt.Sprintf("SELECT ts.id, ts.title, ts.description, ts.status, ts.created_at, ts.updated_at FROM %s ts", tasksTable)

	err := r.db.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskPostgres) DeleteTask(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tasksTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskPostgres) UpdateTask(id int, task taskmanagement.UpdateTaskInput) error {

	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if task.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *task.Title)
		argId++
	}

	if task.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, *task.Description)
		argId++
	}

	if task.Status != nil {
		setValue = append(setValue, fmt.Sprintf("status=$%d", argId))
		args = append(args, *task.Status)
		argId++
	}

	set := strings.Join(setValue, ",")
	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", tasksTable, set, argId)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
