package service

import (
	taskmanagement "taskmanager"
	"taskmanager/pkg/repository"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task taskmanagement.Tasks) (int, error) {
	return s.repo.CreateTask(task)
}
