package controllers

import (
	"bentol/models"
	"bentol/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStores(c *gin.Context) {
	stores, err := services.GetStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreMenus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	store, menues, err := services.GetStoreMenus(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"store": store, "menues": menues})
}

func RegisterStore(c *gin.Context) {
	var input struct {
		Name     string `json:"store_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vendor, err := services.RegisterStore(input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 初回作成時にdefaultのschedule policyを設定
	if err := services.CreateDefaultSchedule(vendor.StoreID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store registered successfully", "vendor": vendor})
}

func LoginStore(c *gin.Context) {
	var input struct {
		Name     string `json:"store_name"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vendor, err := services.LoginStore(input.Name, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "vendor": vendor})
}

func UpdateStorePolicy(c *gin.Context) {
	var input models.StoreBasicReservationPolicy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateStorePolicy(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store policy updated successfully"})
}

func SetSpecificPolicy(c *gin.Context) {
	var input models.StoreSpecificReservationPolicy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.SetSpecificPolicy(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Specific policy set successfully"})
}

func CheckStoreReservation(c *gin.Context) {
}
