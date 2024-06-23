package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "bentol/services"
)

func MakeReservation(c *gin.Context) {
    var input struct {
        UserID   uint   `json:"user_id"`
        StoreID  uint   `json:"store_id"`
        MenueID  uint   `json:"menue_id"`
        Time     string `json:"time"`
        Date     string `json:"date"`
        Count    uint   `json:"count"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    reservTime := input.Date + " " + input.Time
    reservation, err := services.MakeReservation(input.UserID, input.StoreID, input.MenueID, reservTime, input.Count)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Reservation successful", "reservation": reservation})
}

