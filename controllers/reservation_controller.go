package controllers

import (
	"bentol/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeReservation(c *gin.Context) {
	var input struct {
		UserID  uint   `json:"user_id"`
		StoreID uint   `json:"store_id"`
		MenueID uint   `json:"menue_id"`
		Time    string `json:"time"`
		Date    string `json:"date"`
		Count   uint   `json:"count"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservTime := input.Date + " " + input.Time
	reservation, err := services.MakeReservation(input.UserID, input.StoreID, input.MenueID, reservTime, input.Count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation successful", "reservation": reservation})
}

func PayPayPay(c *gin.Context) {
	var payment struct {
		UserID      uint   `json:"user_id"`
		StoreID     uint   `json:"store_id"`
		MenueID     uint   `json:"menue_id"`
		ReservTime  string `json:"reserv_time"`
		ReservCnt   uint   `json:"reserv_cnt"`
		IsRecipt    bool   `json:"is_recipt"`
		TotalAmount uint   `json:"total_amount"`
	}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.ProcessPayment(payment.UserID, payment.StoreID, payment.MenueID, payment.ReservTime, payment.ReservCnt, payment.IsRecipt, payment.TotalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}
