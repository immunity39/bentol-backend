package models

import (
	"time"
)

type UserReservation struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"index"`
	StoreID    uint `gorm:"index"`
	MenueID    uint `gorm:"index"`
	ReservTime time.Time
	ReservCnt  int
	IsReceipt  bool
	Date       time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
