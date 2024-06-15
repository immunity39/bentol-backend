package routes

import (
	"bentol/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/registration", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	r.GET("/store", controllers.GetStores)
	r.GET("/store/:id", controllers.GetStoreMenues)

	r.POST("/store/policy", controllers.SetStorePolicy)

	r.POST("/payment", controllers.CheckAndMakeReservation)
}
