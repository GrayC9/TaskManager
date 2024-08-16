package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Requestor   string    `json:"requestor"`
	Description string    `json:"description"`
	Deadline    string    `json:"deadline"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskStatus struct {
	Status string `json:"status"`
}
