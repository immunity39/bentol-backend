package services

import (
	"bentol/config"
	"bentol/models"
	"bentol/pay"
	"errors"
	"log"
	"time"
)

func MakeReservation(userID, storeID, menueID uint, resDate, resTime string, count uint) (models.UserReservation, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return models.UserReservation{}, err
	}

	reservationTime, err := time.ParseInLocation("2006-01-02 15:04", resDate+" "+resTime, jst)
	if err != nil {
		return models.UserReservation{}, err
	}
	log.Println(reservationTime)

	menue := models.Menue{}
	if err := config.DB.Find(&menue, menueID).Error; err != nil {
		return models.UserReservation{}, err
	}

	reservation := models.UserReservation{
		UserID:      userID,
		StoreID:     storeID,
		MenueID:     menueID,
		ReservTime:  reservationTime,
		ReservCnt:   count,
		IsRecipt:    false,
		TotalAmount: menue.Price * count,
	}

	// 予約可能かどうかの確認
	// スロット毎の最大予約数を取得
	var policy models.StoreSchedule
	config.DB.Where("store_id = ? AND date = ? AND time = ?", storeID, resDate, resTime).First(&policy)

	if int64(count+uint(policy.CurrentReservations)) > int64(policy.MaxReservations) {
		return models.UserReservation{}, errors.New("reservation limit exceeded")
	}

	if err := config.DB.Create(&reservation).Error; err != nil {
		return models.UserReservation{}, errors.New("create user reservation missed")
	}

	if err := config.DB.Model(&models.StoreSchedule{}).Where("id = ?", policy.ID).Update("current_reservations", count+uint(policy.CurrentReservations)).Error; err != nil {
		return models.UserReservation{}, errors.New("update store schedule missed")
	}

	return reservation, nil
}

func ProcessPayment(ReservID, UserID, StoreID, MenueID uint, ReservTime string, ReservCnt uint) (string, error) {

	menue := models.Menue{}
	if err := config.DB.Find(&menue, MenueID).Error; err != nil {
		return "", err
	}

	TotalAmount := menue.Price * ReservCnt

	url, err := pay.Pay(ReservID, menue.Name, TotalAmount)
	if err != nil {
		return "", err
	}

	return url, nil
}

func CancelReservation(ReservID uint) error {
	jst, er := time.LoadLocation("Asia/Tokyo")
	if er != nil {
		return er
	}

	reservation := models.UserReservation{}
	if err := config.DB.Find(&reservation, ReservID).Error; err != nil {
		return err
	}
	log.Println(reservation)

	if reservation.IsRecipt {
		return errors.New("already receipted")
	}

	var err error
	schedule := models.StoreSchedule{}
	err = config.DB.Model(&models.StoreSchedule{}).Where("store_id = ? AND date = ? AND time = ?", reservation.StoreID, reservation.ReservTime.In(jst).Format("2006-01-02"), reservation.ReservTime.In(jst).Format("15:04")).First(&schedule).Error
	if err != nil {
		return err
	}

	err = config.DB.Model(&models.StoreSchedule{}).Where("id = ?", schedule.ID).Update("current_reservations", schedule.CurrentReservations-int(reservation.ReservCnt)).Error
	if err != nil {
		return err
	}

	if err := config.DB.Delete(&reservation).Error; err != nil {
		return err
	}

	return nil
}

func RefundPayment(ReservID uint) error {
	reservation := models.UserReservation{}
	if err := config.DB.Find(&reservation, ReservID).Error; err != nil {
		return err
	}

	totalAmount := reservation.TotalAmount

	if err := pay.Refund(ReservID, totalAmount); err != nil {
		return err
	}

	return nil
}
