package main

import (
	"bentol/config"
	"bentol/cron"
	"bentol/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.SetupRouter(r)

	// Cronジョブの開始
	go cron.StartCronJobs()

	r.Run(":8080")
}
