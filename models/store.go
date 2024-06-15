package models

import (
	"time"
)

type Store struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
