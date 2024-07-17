package controllers

import (
	"bentol/models"
	"bentol/services"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
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

	menues, err := services.GetStoreMenus(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menues": menues})
}

func RegisterStore(c *gin.Context) {
	var input struct {
		Name     string `json:"store_name"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vendor, err := services.RegisterStore(input.Name, input.Mail, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 初回作成時にdefaultのschedule policyを設定
	if err := services.CreateDefaultSchedule(vendor.StoreID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	setStoreSession(c, vendor.StoreID)

	c.JSON(http.StatusOK, gin.H{"message": "Store registered successfully", "vendor": vendor})
}

func LoginStore(c *gin.Context) {
	var input struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vendor, err := services.LoginStore(input.Mail, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	setStoreSession(c, vendor.StoreID)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "vendor": vendor})
}

func setStoreSession(c *gin.Context, storeID uint) {
	session := sessions.Default(c)

	session.Set("storeID", storeID)
	session.Set("role", "store")
	session.Save()
}

func StoreAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		storeID := session.Get("storeID")
		role := session.Get("role")
		if storeID == nil || role != "store" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func LogoutStore(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
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
	var store_id models.Store
	if err := c.ShouldBindJSON(&store_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation, err := services.UpdateCheckStoreReservation(store_id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Check store reservation successfully", "reservation": reservation})
}

func ShipReservation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reservation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	if err := services.ShipReservation(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation shipped successfully"})
}
