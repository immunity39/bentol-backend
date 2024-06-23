package models

import (
    "time"
)

type StoreSchedule struct {
    ID                  uint      `gorm:"primaryKey"`
    StoreID             uint      `gorm:"not null"`
    Date                string    `gorm:"not null"`
    Time                string    `gorm:"not null"`
    MaxReservations     int       `gorm:"not null"`
    CurrentReservations int       `gorm:"not null"`
    CreatedAt           time.Time
    UpdatedAt           time.Time
}
