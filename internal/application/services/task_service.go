package services

import (
	"context"
	"time"

	"github.com/reversersed/taskservice/internal/application/command"
	"github.com/reversersed/taskservice/internal/application/interfaces"
	"github.com/reversersed/taskservice/internal/application/mapper"
	"github.com/reversersed/taskservice/internal/application/query"
	"github.com/reversersed/taskservice/internal/domain/entities"
	"github.com/reversersed/taskservice/internal/domain/repository"
)

type service struct {
	repository repository.TaskRepository
	log        interfaces.Logger
}

func NewTaskService(rep repository.TaskRepository, log interfaces.Logger) interfaces.TaskService {
	return &service{repository: rep, log: log}
}

func (s *service) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	s.log.Infof("received delete request on task %d", id)
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) Update(ctx context.Context, cmd *command.UpdateTaskCommand) (*command.UpdateTaskCommandResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	s.log.Infof("received update command on task %d", cmd.Id)
	task, err := s.repository.Update(ctx, entities.Task{Id: cmd.Id, Title: cmd.Title, Description: cmd.Description, Due: cmd.Due})
	if err != nil {
		return nil, err
	}
	s.log.Infof("task %d has been updated", cmd.Id)
	return &command.UpdateTaskCommandResult{Result: task}, nil
}
func (s *service) Create(ctx context.Context, cmd *command.CreateTaskCommand) (*command.CreateTaskCommandResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	s.log.Infof("received create command on task")
	task, err := s.repository.Create(ctx, cmd.Title, cmd.Description, cmd.Due)
	if err != nil {
		return nil, err
	}
	s.log.Infof("task %s has been created with id %d", task.Title, task.Id)
	return &command.CreateTaskCommandResult{Result: task}, nil
}
func (s *service) GetAll(ctx context.Context) (*query.TaskQueryResultList, error) {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	task, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	s.log.Info("received get all tasks request")
	return mapper.FromListEntityToResult(task), nil
}
func (s *service) Get(ctx context.Context, id int) (*query.TaskQueryResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	s.log.Infof("received get request on task %d", id)
	task, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntityToResult(task), nil
}
