package request

import (
	"time"

	"github.com/reversersed/taskservice/internal/application/command"
)

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	DueTime     string `json:"due" validate:"required"`
}

func (c *CreateTaskRequest) Command() (*command.CreateTaskCommand, error) {
	time, err := time.Parse("2006-01-02T15:04:05", c.DueTime)
	if err != nil {
		return nil, err
	}
	return &command.CreateTaskCommand{Title: c.Title, Description: c.Description, Due: time}, nil
}
