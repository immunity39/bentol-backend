package controllers

import (
	"bentol/config"
	"bentol/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RegistrationInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var input RegistrationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Password: input.Password, Mail: input.Mail, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

type LoginInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("name = ? AND password = ?", input.Name, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
