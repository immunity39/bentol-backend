package controllers

import (
	"bentol/config"
	"bentol/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAndMakeReservation(c *gin.Context) {
	var input models.PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservTime, err := time.Parse("15:04", input.Time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}

	reservDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	var policy models.StoreReservationPolicy
	if err := config.DB.Where("store_id = ? AND (date = ? OR (date IS NULL AND day_of_week = ?))",
		input.StoreID, reservDate, int(reservDate.Weekday())).First(&policy).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation policy not found"})
		return
	}

	slotTime := reservTime.Truncate(time.Duration(policy.TimeSlotInterval) * time.Minute)

	var timeSlot models.StoreTimeSlotReservation
	config.DB.Where("store_id = ? AND date = ? AND time_slot = ?", input.StoreID, reservDate, slotTime).FirstOrCreate(&timeSlot, models.StoreTimeSlotReservation{
		StoreID:             input.StoreID,
		Date:                reservDate,
		TimeSlot:            slotTime,
		CurrentReservations: 0,
	})

	if timeSlot.CurrentReservations+input.Count > policy.MaxReservationsPerSlot {
		c.JSON(http.StatusConflict, gin.H{"error": "Reservation limit exceeded"})
		return
	}

	timeSlot.CurrentReservations += input.Count
	config.DB.Save(&timeSlot)

	userReservation := models.UserReservation{
		UserID:     input.UserID,
		StoreID:    input.StoreID,
		MenueID:    input.MenueID,
		ReservTime: reservTime,
		ReservCnt:  input.Count,
		IsReceipt:  false,
		Date:       reservDate,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	config.DB.Create(&userReservation)

	c.JSON(http.StatusOK, gin.H{"message": "Reservation successful"})
}
