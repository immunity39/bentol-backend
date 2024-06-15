package models

import (
	"time"
)

type StoreTimeSlotReservation struct {
	ID                  uint `gorm:"primaryKey"`
	StoreID             uint `gorm:"index"`
	Date                time.Time
	TimeSlot            time.Time
	CurrentReservations int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
