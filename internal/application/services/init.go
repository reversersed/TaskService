package services

import (
	"github.com/reversersed/taskservice/internal/application/interfaces"
	"github.com/reversersed/taskservice/internal/domain/repository"
)

type service struct {
	repository repository.TaskRepository
	log        interfaces.Logger
}

func New(rep repository.TaskRepository, log interfaces.Logger) *service {
	return &service{repository: rep, log: log}
}
