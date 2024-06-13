package handler

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/handler/validator"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := repository.FindUserByName(user.Name)
	if err != nil || !usecase.CheckPasswordHash(user.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failure", "message": "Invalid credentials"})
		return
	}

	token, err := usecase.GenerateJWT(dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failure", "message": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}

func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.ValidateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := usecase.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	if err := repository.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User registered successfully"})
}

func GetStores(c *gin.Context) {
	stores, err := repository.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreMenues(c *gin.Context) {
	storeID := c.Param("id")
	menues, err := repository.GetStoreMenues(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"menues": menues})
}

func GetMenueAvailability(c *gin.Context) {
	menueID := c.Param("id")
	availability, err := repository.GetMenueAvailability(menueID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"availability": availability})
}

func MakePayment(c *gin.Context) {
	var reservation model.UserReservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check availability logic should be here

	if err := repository.MakeReservation(reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Reservation made successfully"})
}

