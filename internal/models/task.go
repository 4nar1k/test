package models

import "gorm.io/gorm"

// Task - основная модель для хранения задач
type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"` // Связь с каскадным удалением
}
