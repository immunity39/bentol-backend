package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100)"`
	Mail      string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
