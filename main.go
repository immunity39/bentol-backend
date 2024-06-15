package main

import (
	"bentol/config"
	"bentol/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
