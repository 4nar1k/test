package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	
	// Связь с задачами
	Tasks []Task `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
