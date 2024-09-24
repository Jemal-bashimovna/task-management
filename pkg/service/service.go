package service

import (
	taskmanagement "taskmanager"
	"taskmanager/pkg/repository"
)

type Tasks interface {
	CreateTask(task taskmanagement.Tasks) (int, error)
	GetTasks() ([]taskmanagement.Tasks, error)
	DeleteTask(id int) error
}

type Service struct {
	Tasks
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Tasks: NewTaskService(repo.Tasks),
	}
}
