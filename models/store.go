package models

import "time"

type Store struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"type:varchar(100)"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
