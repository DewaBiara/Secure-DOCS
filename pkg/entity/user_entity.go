package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;"`
	Username  string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password  string `gorm:"type:varchar(255);not null"`
	Name      string
	Telp      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Users []User
