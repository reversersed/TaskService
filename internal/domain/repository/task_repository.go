package repository

import (
	"context"
	"time"

	"github.com/reversersed/taskservice/internal/domain/entities"
)

type TaskRepository interface {
	Create(context.Context, string, string, time.Time) (entities.Task, error)
	Update(context.Context, entities.Task) (entities.Task, error)
	Delete(context.Context, int) error
	GetAll(context.Context) ([]entities.Task, error)
	Get(context.Context, int) (entities.Task, error)
}
