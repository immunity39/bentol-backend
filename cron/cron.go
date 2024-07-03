package cron

import (
	"bentol/services"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func StartCronJobs() {
	tky, _ := time.LoadLocation("Asia/Tokyo")
	c := cron.New(cron.WithLocation(tky))

	// 毎日午前3時に実行
	c.AddFunc("0 3 * * *", func() {
		log.Println("Running create weekly schedules cron job")
		if err := services.CreateWeeklySchedules(); err != nil {
			log.Println("Error creating weekly schedules:", err)
		}
	})

	// 毎日午前3:30時に実行
	c.AddFunc("30 3 * * *", func() {
		log.Println("Running update specific schedules cron job")
		if err := services.UpdateSpecificSchedules(); err != nil {
			log.Println("Error updating specific schedules:", err)
		}
	})

	c.Start()
}
