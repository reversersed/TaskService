package entities

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due         time.Time `json:"due_date"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}
