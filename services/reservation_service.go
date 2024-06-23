package services

import (
    "bentol/config"
    "bentol/models"
    "time"
    "errors"
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
