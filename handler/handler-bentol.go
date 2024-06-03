package handler

import (
	"bentol/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUsecase *usecase.LoginUsecase
}

func NewLoginHandler(lu *usecase.LoginUsecase) *LoginHandler {
	return &LoginHandler{LoginUsecase: lu}
}

func (lh *LoginHandler) Login(c *gin.Context) {
	var loginData struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	user, err := lh.LoginUsecase.Login(loginData.Name, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
