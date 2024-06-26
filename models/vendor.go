package models

import "time"

type StoreVendor struct {
	ID        uint `gorm:"primaryKey"`
	StoreID   uint
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
