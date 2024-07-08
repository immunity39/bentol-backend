package main

import (
	"bentol/config"
	"bentol/cron"
	"bentol/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://133.14.14.14:8090",
			"https://133.14.14.14:8090",
			"http://bentol.rd.dendai.ac.jp",
			"https://bentol.rd.dendai.ac.jp",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	routes.SetupRouter(r)

	// Cronジョブの開始
	go cron.StartCronJobs()

	r.Run(":8090")
}
