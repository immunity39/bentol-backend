package models

import (
	"time"
)

type StoreBasicReservationPolicy struct {
	ID                     uint   `gorm:"primaryKey"`
	StoreID                uint   `gorm:"not null"`
	DayOfWeek              int    `gorm:"not null"`
	TimeSlotInterval       int    `gorm:"not null"`
	MaxReservationsPerSlot int    `gorm:"not null"`
	StoreStartTime         string `gorm:"not null"`
	StoreEndTime           string `gorm:"not null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

type StoreSpecificReservationPolicy struct {
	ID                     uint   `gorm:"primaryKey"`
	StoreID                uint   `gorm:"not null"`
	Date                   string `gorm:"not null"`
	TimeSlotInterval       int    `gorm:"not null"`
	MaxReservationsPerSlot int    `gorm:"not null"`
	StoreStartTime         string `gorm:"not null"`
	StoreEndTime           string `gorm:"not null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
