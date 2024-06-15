package models

import (
	"time"
)

type StoreReservationPolicy struct {
	ID                     uint `gorm:"primaryKey"`
	StoreID                uint `gorm:"index"`
	DayOfWeek              int
	Date                   *time.Time
	TimeSlotInterval       int
	MaxReservationsPerSlot int
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
