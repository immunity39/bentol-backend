package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Mail      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
