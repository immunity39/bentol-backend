package services

import (
	"bentol/config"
	"bentol/models"
	"errors"
)

func GetStores() ([]models.Store, error) {
	var stores []models.Store
	if err := config.DB.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStoreMenus(storeID uint) (models.Store, []models.Menue, error) {
	var store models.Store
	var menues []models.Menue

	if err := config.DB.First(&store, storeID).Error; err != nil {
		return models.Store{}, nil, errors.New("store not found")
	}
	if err := config.DB.Where("store_id = ?", storeID).Find(&menues).Error; err != nil {
		return store, nil, err
	}
	return store, menues, nil
}

func RegisterStore(name, email, password string) (models.StoreVendor, error) {
	store := models.Store{Name: name}
	if err := config.DB.Create(&store).Error; err != nil {
		return models.StoreVendor{}, err
	}

	vendor := models.StoreVendor{StoreID: store.ID, Name: name, Email: email, Password: password}
	if err := config.DB.Create(&vendor).Error; err != nil {
		return models.StoreVendor{}, err
	}

	return vendor, nil
}

func CreateDefaultSchedule(sid uint) error {
	for day := 0; day < 7; day++ {
		policy := models.StoreBasicReservationPolicy{
			StoreID:                sid,
			DayOfWeek:              day,
			TimeSlotInterval:       10,
			MaxReservationsPerSlot: 10,
			StoreStartTime:         "12:50",
			StoreEndTime:           "13:40",
		}

		if err := config.DB.Create(&policy).Error; err != nil {
			return err
		}
	}

	return nil
}

func LoginStore(name, password string) (models.StoreVendor, error) {
	var vendor models.StoreVendor
	if err := config.DB.Where("name = ? AND password = ?", name, password).First(&vendor).Error; err != nil {
		return models.StoreVendor{}, errors.New("invalid credentials")
	}
	return vendor, nil
}

func UpdateStorePolicy(policy models.StoreBasicReservationPolicy) error {
	if err := config.DB.Model(&models.StoreBasicReservationPolicy{}).Where("store_id = ?", policy.StoreID).Updates(policy).Error; err != nil {
		return err
	}
	return nil
}

func SetSpecificPolicy(policy models.StoreSpecificReservationPolicy) error {
	if err := config.DB.Create(&policy).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCheckStoreReservation() error {
	/*
		必要な実装
		storeが営業時間の時にcron jobが動作してほしい。
		store_idごとに予約をわけて実装するべきでは？ -> 呼び出し時に store_idを渡す方針のほうがよさそう
		取得するuser reservationは store_id = id and now_time < reseervation_time and 受け取りフラグ != trueのデータ
	*/
	return nil
}
