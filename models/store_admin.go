package models

import (
	"time"
)

type StoreAdmin struct {
	ID        uint `gorm:"primaryKey"`
	StoreID   uint
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
