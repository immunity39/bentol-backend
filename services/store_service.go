package services

import (
	"bentol/config"
	"bentol/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetStores() ([]models.Store, error) {
	var stores []models.Store
	if err := config.DB.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStoreMenus(storeID uint) ([]models.Menue, error) {
	var store models.Store
	var menues []models.Menue

	if err := config.DB.First(&store, storeID).Error; err != nil {
		return nil, errors.New("store not found")
	}
	if err := config.DB.Where("store_id = ?", storeID).Find(&menues).Error; err != nil {
		return nil, err
	}
	return menues, nil
}

func RegisterStore(name, mail, password string) (models.StoreVendor, error) {
	store := models.Store{Name: name}
	if err := config.DB.Create(&store).Error; err != nil {
		return models.StoreVendor{}, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.StoreVendor{}, err
	}

	vendor := models.StoreVendor{StoreID: store.ID, Name: name, Mail: mail, Password: string(hashPassword)}
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

func LoginStore(mail, password string) (models.StoreVendor, error) {
	var vendor models.StoreVendor
	if err := config.DB.Where("mail = ?", mail).First(&vendor).Error; err != nil {
		return models.StoreVendor{}, errors.New("vendor not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(vendor.Password), []byte(password)); err != nil {
		return models.StoreVendor{}, errors.New("incorrect password")
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
	var resipt = false
	if err := config.DB.Where("store_id = ? AND is_recipt = ?", store, resipt).Find(&reservation).Error; err != nil {
		return []models.UserReservation{}, err
	}
	return reservation, nil
}

func ShipReservation(reservationID uint) error {
	var reservation models.UserReservation
	if err := config.DB.Where("id = ?", reservationID).First(&reservation).Error; err != nil {
		return errors.New("reservation not found")
	}
	reservation.IsRecipt = true
	if err := config.DB.Save(&reservation).Error; err != nil {
		return err
	}
	return nil
}
