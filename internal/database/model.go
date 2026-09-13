package database

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
}
