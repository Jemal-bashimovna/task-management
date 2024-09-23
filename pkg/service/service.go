package service

import "taskmanager/pkg/repository"

type Tasks interface {
}

type Service struct {
	Tasks
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
