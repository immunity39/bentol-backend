package services

import (
	"bentol/config"
	"bentol/models"
	"errors"
	"time"
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

func UpdateCheckStoreReservation(store uint) ([]models.UserReservation, error) {
	var reservation []models.UserReservation
	var now_time = time.Now()
	var resipt = false
	if err := config.DB.Where("store_id = ? AND reserv_time > ? AND is_recipt = ?", store, now_time, resipt).Find(&reservation).Error; err != nil {
		return []models.UserReservation{}, err
	}
	return reservation, nil
}
