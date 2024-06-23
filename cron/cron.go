package cron

import (
	"bentol/services"
	"log"

	"github.com/robfig/cron/v3"
)

func StartCronJobs() {
	c := cron.New(cron.WithSeconds())

	// 毎日午前3時に実行
	c.AddFunc("0 0 3 * * *", func() {
		log.Println("Running create weekly schedules cron job")
		if err := services.CreateWeeklySchedules(); err != nil {
			log.Println("Error creating weekly schedules:", err)
		}
	})

	// 毎日午前4時に実行
	c.AddFunc("0 0 4 * * *", func() {
		log.Println("Running update specific schedules cron job")
		if err := services.UpdateSpecificSchedules(); err != nil {
			log.Println("Error updating specific schedules:", err)
		}
	})

	c.Start()

	// 永遠に実行されるように待機
	select {}
}
