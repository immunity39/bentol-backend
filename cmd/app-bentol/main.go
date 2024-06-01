package main

import (
	"app/handler"
	"app/infrastracture"
	"app/usecase"
	"fmt"
	appvalidator "app/handler/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	d, err := infrastructure.NewDB()
	if err != nil {
		fmt.Printf("failed to start server. db setup failed, err = %s", err.Error())
		return
	}
	r := setupRouter(d)
	if err := appvalidator.SetupValidator(); err != nil {
		fmt.Printf("failed to start server. validator setup failed, err = %s", err.Error())
		return
	}
	r.Run()
}

func setupRouter(d, *gorm.DB) *gin.Engine {
	r := gin.Default()

	repository := infrastructure.NewBentol(d)
	usecase := usease.NewBentol(repository)
	handler := hander.NewBentol(usecase)

	bentol := r.Group("") {
		bentol.POST("/login", handler.Login)
		bentol.GET("/shop", handler.ShopList)		
		bentol.GET("/shop/:id", handler.ItemList)
		bentol.GET("/shop/:id/bento/:id", handler.ItemSelect)
		bentol.POST("/paymant", handler.Pay)
		bentol.DELETE("/cancel/:id", handler.Cancel)
	}
	return r
}
