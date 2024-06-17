package models

import (
	"time"
)

type StoreManager struct {
	ID        uint   `gorm:"primaryKey"`
	StoreID   uint   `gorm:"not null"`
	Password  string `gorm:"not null"`
	Mail      string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
