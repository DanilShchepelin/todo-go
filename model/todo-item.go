package model

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
}
