package main

import (
	"backend/handler"
	"backend/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	infrastructure.InitDB()

	router.POST("/login", handler.Login)
	router.POST("/registration", handler.RegisterUser)
	router.GET("/store", handler.GetStores)
	router.GET("/store/:id", handler.GetStoreMenues)
	router.GET("/menue/:id", handler.GetMenueAvailability)
	router.POST("/payment", handler.MakePayment)

	router.Run(":8080")
}
