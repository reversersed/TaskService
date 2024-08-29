package repository

import "time"

type Task struct {
	Id          int       `db:"Id"`
	Title       string    `db:"Title"`
	Description string    `db:"Description"`
	Due         time.Time `db:"Due"`
	Created     time.Time `db:"Created"`
	Updated     time.Time `db:"Updated"`
}
