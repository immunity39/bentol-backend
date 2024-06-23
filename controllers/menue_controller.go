package controllers

import (
	"bentol/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMenue(c *gin.Context) {
	var input struct {
		StoreID     uint   `json:"store_id"`
		Name        string `json:"name"`
		Price       uint   `json:"price"`
		Description string `json:"description"`
		IsSoldOut   bool   `json:"is_sold_out"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menue, err := services.AddMenue(input.StoreID, input.Name, input.Price, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menue added successfully", "menue": menue})
}

func UpdateMenue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menue ID"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Price       uint   `json:"price"`
		Description string `json:"description"`
		IsSoldOut   bool   `json:"is_sold_out"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menue, err := services.UpdateMenue(uint(id), input.Name, input.Price, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menue updated successfully", "menue": menue})
}
