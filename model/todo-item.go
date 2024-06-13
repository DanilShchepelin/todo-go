package model

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

type TodoItem struct {
	ID          uint   `gorm:"primarykey"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
