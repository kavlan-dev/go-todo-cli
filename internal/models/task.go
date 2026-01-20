package models

import "time"

type Task struct {
	Id          int        `json:"id"`
	Content     string     `json:"content"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
