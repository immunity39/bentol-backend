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

	reservation, err := services.MakeReservation(input.UserID, input.StoreID, input.MenueID, input.Date, input.Time, input.Count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservTime := input.Date + " " + input.Time
	url, err := services.ProcessPayment(reservation.ID, input.UserID, input.StoreID, input.MenueID, reservTime, input.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = services.CurrentReservationUpdate(reservation.ID, input.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation successful", "URL": url})
}

func CancelReservation(c *gin.Context) {
	var input struct {
		ReservID uint `json:"reservation_id"`
	}
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.RefundPayment(input.ReservID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.CancelReservation(input.ReservID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = services.ReservationDelete(input.ReservID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Refund successful"})
}
