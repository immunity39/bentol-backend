package controllers

import (
	"bentol/config"
	"bentol/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStores(c *gin.Context) {
	var stores []models.Store
	config.DB.Find(&stores)
	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreMenues(c *gin.Context) {
	var menues []models.Menue
	storeID := c.Param("id")
	config.DB.Where("store_id = ?", storeID).Find(&menues)
	c.JSON(http.StatusOK, gin.H{"menues": menues})
}

type PolicyInput struct {
	StoreID                uint    `json:"store_id" binding:"required"`
	DayOfWeek              int     `json:"day_of_week"`
	Date                   *string `json:"date"`
	TimeSlotInterval       int     `json:"time_slot_interval" binding:"required"`
	MaxReservationsPerSlot int     `json:"max_reservations_per_slot" binding:"required"`
}

func SetStorePolicy(c *gin.Context) {
	var input PolicyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var date *time.Time
	if input.Date != nil {
		parsedDate, err := time.Parse("2006-01-02", *input.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		date = &parsedDate
	}

	policy := models.StoreReservationPolicy{
		StoreID:                input.StoreID,
		DayOfWeek:              input.DayOfWeek,
		Date:                   date,
		TimeSlotInterval:       input.TimeSlotInterval,
		MaxReservationsPerSlot: input.MaxReservationsPerSlot,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	result := config.DB.Create(&policy)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Policy set successfully"})
}
