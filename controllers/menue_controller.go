package controllers

import (
	"bentol/config"
	"bentol/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMenue(c *gin.Context) {
	var input struct {
		StoreID     uint   `json:"store_id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		Price       int    `json:"price" binding:"required"`
		Description string `json:"description" binding:"required"`
		IsSoldOut   bool   `json:"is_sold_out"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menue := models.Menue{
		StoreID:     input.StoreID,
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		IsSoldOut:   input.IsSoldOut,
	}

	if err := config.DB.Create(&menue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menue added successfully"})
}

func UpdateMenue(c *gin.Context) {
	menueID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menue ID"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		IsSoldOut   bool   `json:"is_sold_out"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var menue models.Menue
	if err := config.DB.First(&menue, menueID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menue not found"})
		return
	}

	menue.Name = input.Name
	menue.Price = input.Price
	menue.Description = input.Description
	menue.IsSoldOut = input.IsSoldOut

	if err := config.DB.Save(&menue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menue updated successfully"})
}
