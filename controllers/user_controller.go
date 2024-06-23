package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "bentol/services"
)

func RegisterUser(c *gin.Context) {
    var input struct {
        Name     string `json:"name"`
        Password string `json:"password"`
        Mail     string `json:"mail"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := services.RegisterUser(input.Name, input.Password, input.Mail)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}

func LoginUser(c *gin.Context) {
    var input struct {
        Name     string `json:"name"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := services.LoginUser(input.Name, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

