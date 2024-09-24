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
func (s *TaskService) GetTasks() ([]taskmanagement.Tasks, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) UpdateTask(id int, task taskmanagement.UpdateTaskInput) error {
	if err := task.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateTask(id, task)
}
