package entities

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	Due         time.Time
	Created     time.Time
	Updated     time.Time
}
