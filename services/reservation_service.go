package services

import (
	"bentol/config"
	"bentol/models"
	"encoding/json"
	"errors"
	"net/http"
	"os/exec"
	"time"
)

func MakeReservation(userID, storeID, menueID uint, reservTime string, count uint) (models.UserReservation, error) {
	reservationTime, err := time.Parse("2006-01-02 15:04", reservTime)
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
	var existingReservations int64
	config.DB.Model(&models.UserReservation{}).Where("store_id = ? AND menue_id = ? AND reserv_time = ?", storeID, menueID, reservationTime).Count(&existingReservations)

	// スロット毎の最大予約数を取得
	var policy models.StoreBasicReservationPolicy
	config.DB.Where("store_id = ? AND day_of_week = DAYOFWEEK(?)", storeID, reservationTime).First(&policy)

	if existingReservations >= int64(policy.MaxReservationsPerSlot) {
		return models.UserReservation{}, errors.New("reservation limit exceeded")
	}

	if err := config.DB.Create(&reservation).Error; err != nil {
		return models.UserReservation{}, err
	}
	return reservation, nil
}

// remake
/*
get menue
totalamount = count * menue.price
description = menue.name

pay/payment.goを呼び出す + 支払いのuuid を保持する必要があればDBの形式を変えること
*/
func ProcessPayment(ReservID, UserID, StoreID, MenueID uint, ReservTime string, ReservCnt uint, IsRecipt bool, TotalAmount uint) error {
	var PaymentRequest struct {
		ReservID    uint    `json:"reserv_id"`
		UserID      uint    `json:"user_id"`
		StoreID     uint    `json:"store_id"`
		MenueID     uint    `json:"menue_id"`
		ReservTime  string  `json:"reserv_time"`
		ReservCnt   uint    `json:"reserv_cnt"`
		IsRecipt    bool    `json:"is_recipt"`
		TotalAmount float64 `json:"total_amount"`
	}
	requestBody, err := json.Marshal(PaymentRequest)
	if err != nil {
		return err
	}

	resp, err := exec.Command("node ./pay/paypayapi.js", string(requestBody)).Output()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to process payment")
	}
	return nil
}
