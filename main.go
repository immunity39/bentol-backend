package main

import (
	"bentol/config"
	"bentol/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	routes.SetupRouter()

	r.Run(":8080")
}
