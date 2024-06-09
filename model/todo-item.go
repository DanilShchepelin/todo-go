package model

import (
	_ "gorm.io/gorm"
	"time"
)

type TodoItem struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
}
