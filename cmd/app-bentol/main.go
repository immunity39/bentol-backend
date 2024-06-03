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

	r := gin.Default()
	r.POST("/login", loginHandler.Login)
	r.Run(":8080")
}
