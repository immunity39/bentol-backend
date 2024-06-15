package models

type PaymentInput struct {
	UserID  uint   `json:"user_id" binding:"required"`
	StoreID uint   `json:"store_id" binding:"required"`
	MenueID uint   `json:"menue_id" binding:"required"`
	Time    string `json:"time" binding:"required"`
	Date    string `json:"date" binding:"required"`
	Count   int    `json:"count" binding:"required"`
}
