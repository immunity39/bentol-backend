             int    `json:"id"`
	StoreID        int    `json:"store_id"`
	Weekday        string `json:"weekday"`
	TimeSlot       string `json:"time_slot"`
	MaxReservations int    `json:"max_reservations"`
}
package model

type User struct {
	ID       int    `json:"id"`
	Mail     string `json:"mail"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Store struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Menue struct {
	ID          int    `json:"id"`
	StoreID     int    `json:"store_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	IsSoldOut   bool   `json:"is_sold_out"`
}

type UserReservation struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	StoreID    int    `json:"store_id"`
	MenueID    int    `json:"menue_id"`
	ReservTime string `json:"reserv_time"`
	ReservDate string `json:"reserv_date"`
	ReservCnt  int    `json:"reserv_cnt"`
	IsRecipt   bool   `json:"is_recipt"`
}

type StoreSchedule struct {
	ID
