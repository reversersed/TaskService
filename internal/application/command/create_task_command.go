package command

import "time"

type CreateTaskCommand struct {
	Title       string
	Description string
	Due         time.Time
}
