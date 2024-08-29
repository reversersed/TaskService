package command

import "time"

type UpdateTaskCommand struct {
	Id          int
	Title       string
	Description string
	Due         time.Time
}
