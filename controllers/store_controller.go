package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"bentol/config"
	"bentol/models"
)

func RegisterStore(c *gin.Context) {
	var input struct {
		StoreName string `json:"store_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := models.Store{Name: input.StoreName}
	if err := config.DB.Create(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	storeAdmin := models.StoreAdmin{
		StoreID:  store.ID,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := config.DB.Create(&storeAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store and admin registered successfully"})
}

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

func UpdateStore(c *gin.Context) {
	storeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	var input struct {
		StoreName string `json:"store_name"`
		Policy    struct {
			TimeSlotInterval       int `json:"time_slot_interval"`
			MaxReservationsPerSlot int `json:"max_reservations_per_slot"`
		} `json:"policy"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var store models.Store
	if err := config.DB.First(&store, storeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	store.Name = input.StoreName
	if err := config.DB.Save(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	policy := models.StoreReservationPolicy{
		StoreID:                uint(storeID),
		TimeSlotInterval:       input.Policy.TimeSlotInterval,
		MaxReservationsPerSlot: input.Policy.MaxReservationsPerSlot,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	if err := config.DB.Save(&policy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store and policy updated successfully"})

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
