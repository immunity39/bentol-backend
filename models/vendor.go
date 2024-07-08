package models

import "time"

type StoreVendor struct {
	ID        uint `gorm:"primaryKey"`
	StoreID   uint
	Name      string `gorm:"type:varchar(100)"`
	Mail      string `gorm:"type:varchar(100);unique"`
	Password  string `gorm:"type:varchar(100);unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
