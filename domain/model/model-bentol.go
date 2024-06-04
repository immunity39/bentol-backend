package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"unique"`
	Password string
}

type Store struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type Menue struct {
	ID          uint `gorm:"primaryKey"`
	StoreID     uint
	Name        string
	Price       int
	Description string
	IsSoldOut   bool
}

type UserReservation struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	StoreID    uint
	MenueID    uint
	ReservTime string
	ReservCnt  int
	IsRecipt   bool
}
