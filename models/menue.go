package models

import "time"

type Menue struct {
	ID          uint `gorm:"primaryKey"`
	StoreID     uint
	Name        string `gorm:"type:varchar(100)"`
	Price       uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
