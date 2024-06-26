package models

import "time"

type UserReservation struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	StoreID     uint
	MenueID     uint
	ReservTime  time.Time
	ReservCnt   uint
	IsRecipt    bool
	TotalAmount uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
