package models

import (
	"time"
)

// Task represent the task model
type Task struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Comment   string     `json:"comment"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	UserID    int64      `json:"user_id"`
}
