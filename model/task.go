package model

import (
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Done      bool      `json:"done"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}
