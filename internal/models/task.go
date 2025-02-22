package models

import "gorm.io/gorm"

// Task - основная модель для хранения задач
type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
