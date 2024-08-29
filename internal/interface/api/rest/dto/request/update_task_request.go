package request

import (
	"time"

	"github.com/reversersed/taskservice/internal/application/command"
)

type UpdateTaskRequest struct {
	Id          int    `json:"-"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueTime     string `json:"due" validate:"required"`
}

func (c *UpdateTaskRequest) Command() (*command.UpdateTaskCommand, error) {
	time, err := time.Parse("2006-01-02T15:04:05", c.DueTime)
	if err != nil {
		return nil, err
	}
	return &command.UpdateTaskCommand{Id: c.Id, Title: c.Title, Description: c.Description, Due: time}, nil
}
