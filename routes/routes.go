package routes

import (
	"bentol/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/registration", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Store routes
	router.GET("/store", controllers.GetStores)
	router.POST("/store/register", controllers.RegisterStore)
	router.PUT("/store/:id/update", controllers.UpdateStore)

	// Menu routes
	router.GET("/store/:id", controllers.GetStoreMenues)
	router.POST("/menue/add", controllers.AddMenue)
	router.PUT("/menue/:id/update", controllers.UpdateMenue)

	// Reservation routes
	router.POST("/payment", controllers.CheckAndMakeReservation)

	return router
}
