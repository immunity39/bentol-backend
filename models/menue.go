package models

import (
	"time"
)

type Menue struct {
	ID          uint `gorm:"primaryKey"`
	StoreID     uint `gorm:"index"`
	Name        string
	Price       int
	Description string
	IsSoldOut   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
