package services

import (
	"bentol/config"
	"bentol/models"
	"bentol/pay"
	"errors"
	"time"
)

func MakeReservation(userID, storeID, menueID uint, resDate, resTime string, count uint) (models.UserReservation, error) {
	reservationTime, err := time.Parse("2006-01-02 15:04", resDate+" "+resTime)
	if err != nil {
		return models.UserReservation{}, err
	}

	reservation := models.UserReservation{
		UserID:     userID,
		StoreID:    storeID,
		MenueID:    menueID,
		ReservTime: reservationTime,
		ReservCnt:  count,
		IsRecipt:   false,
	}

	// 予約可能かどうかの確認
	// スロット毎の最大予約数を取得
	var policy models.StoreSchedule
	config.DB.Where("store_id = ? AND date = ? AND time = ?").First(&policy)

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

// pay/payment.goを呼び出す + 支払いのuuid を保持する必要があればDBの形式を変えること
// err = services.ProcessPayment(reservation.ID, input.UserID, input.StoreID, input.MenueID, reservTime, input.Count)
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
