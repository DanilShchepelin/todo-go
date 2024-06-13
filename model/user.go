package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint    `gorm:"primarykey"`
	Name       string  `gorm:"not null" json:"name"`
	LastName   string  `json:"lastName"`
	Patronymic *string `json:"patronymic"`
	FullName   string  `json:"fullName"`
	Password   string  `gorm:"->:false;<-:create" json:"password"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
