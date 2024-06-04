package main

import (
	"bentol/handler"
	"bentol/infrastructure"
	"bentol/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.InitDB()

	userRepository := infrastructure.NewUserRepository()
	loginUsecase := usecase.NewLoginUsecase(userRepository)
	loginHandler := handler.NewLoginHandler(loginUsecase)

	storeRepository := infrastructure.NewStoreRepository()
	storeUsecase := usecase.NewStoreUsecase(storeRepository)
	storeHandler := handler.NewStoreHandler(storeUsecase)

	menueRepository := infrastructure.NewMenueRepository()
	menueUsecase := usecase.NewMenueUsecase(menueRepository)
	menueHandler := handler.NewMenueHandler(menueUsecase)

	userReservationRepository := infrastructure.NewUserReservationRepository()
	userReservationUsecase := usecase.NewUserReservationUsecase(userReservationRepository)
	userReservationHandler := handler.NewUserReservationHandler(userReservationUsecase)

	r := gin.Default()
	r.POST("/login", loginHandler.Login)
	r.GET("/store", storeHandler.GetAllStores)
	r.GET("/store/:id", menueHandler.GetMenuesByStoreID)
	r.GET("/menue/:id", menueHandler.GetMenueByID)
	r.POST("/payment", userReservationHandler.CreateReservation)
	r.DELETE("/cancel/:id", userReservationHandler.DeleteReservation)
	r.Run(":8080")
}
