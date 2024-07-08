package services

import (
	"bentol/config"
	"bentol/models"
	"log"
	"time"
)

func CreateWeeklySchedules() error {
	// 現在の日付を取得
	today := time.Now()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// 1週間先までのスケジュールを作成
	for i := 0; i < 7; i++ {
		date := today.AddDate(0, 0, i).In(jst).Format("2006-01-02")
		weekday := (today.AddDate(0, 0, i+1).Weekday() + 6) % 7 // 月曜日を0とする

		var stores []models.Store
		if err := config.DB.Find(&stores).Error; err != nil {
			return err
		}

		for _, store := range stores {
			var policy models.StoreBasicReservationPolicy
			if err := config.DB.Where("store_id = ? AND day_of_week = ?", store.ID, weekday).First(&policy).Error; err != nil {
				log.Printf("No policy found for store %d on weekday %d\n", store.ID, weekday)
				continue
			}

			startTime, _ := time.ParseInLocation("15:04", policy.StoreStartTime, jst)
			endTime, _ := time.ParseInLocation("15:04", policy.StoreEndTime, jst)
			interval := time.Duration(policy.TimeSlotInterval) * time.Minute

			for t := startTime; t.Before(endTime); t = t.Add(interval) {
				schedule := models.StoreSchedule{
					StoreID:             store.ID,
					Date:                date,
					Time:                t.In(jst).Format("15:04"),
					MaxReservations:     policy.MaxReservationsPerSlot,
					CurrentReservations: 0,
				}
				if err := config.DB.Create(&schedule).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func UpdateSpecificSchedules() error {
	// 現在の日付を取得
	today := time.Now()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// 1週間先までのスケジュールを更新
	for i := 0; i < 7; i++ {
		date := today.AddDate(0, 0, i).In(jst).Format("2006-01-02")

		var specificPolicies []models.StoreSpecificReservationPolicy
		if err := config.DB.Where("date = ?", date).Find(&specificPolicies).Error; err != nil {
			return err
		}

		for _, policy := range specificPolicies {
			startTime, _ := time.ParseInLocation("15:04", policy.StoreStartTime, jst)
			endTime, _ := time.ParseInLocation("15:04", policy.StoreEndTime, jst)
			interval := time.Duration(policy.TimeSlotInterval) * time.Minute

			if err := config.DB.Where("store_id = ? AND date = ?", policy.StoreID, date).Delete(&models.StoreSchedule{}).Error; err != nil {
				return err
			}

			for t := startTime; t.Before(endTime); t = t.Add(interval) {
				schedule := models.StoreSchedule{
					StoreID:             policy.StoreID,
					Date:                date,
					Time:                t.In(jst).Format("15:04"),
					MaxReservations:     policy.MaxReservationsPerSlot,
					CurrentReservations: 0,
				}
				if err := config.DB.Create(&schedule).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
