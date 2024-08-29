package interfaces

import (
	"context"

	"github.com/reversersed/taskservice/internal/application/command"
	"github.com/reversersed/taskservice/internal/application/query"
)

type TaskService interface {
	Delete(context.Context, int) error
	Update(context.Context, *command.UpdateTaskCommand) (*command.UpdateTaskCommandResult, error)
	Create(context.Context, *command.CreateTaskCommand) (*command.CreateTaskCommandResult, error)
	GetAll(context.Context) (*query.TaskQueryResultList, error)
	Get(context.Context, int) (*query.TaskQueryResult, error)
}
