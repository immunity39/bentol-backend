ity(menueID int) ([]model.StoreSchedule, error) {
	var schedules []model.StoreSchedule
	if err := infrastructure.DB.Where("menue_id = ?", menueID).Find(&schedules).Error; err != nil {
		return nil, err
	}
; err != nil {
		return err
	}
	return nil
}
	return schedules, nil
}

func MakeReservation(reservation model.UserReservation) error {
	if err := infrastructure.DB.Create(&reservation).Errorpackage repository

import (
	"backend/domain/model"
	"backend/infrastructure"
)

func FindUserByName(name string) (model.User, error) {
	var user model.User
	if err := infrastructure.DB.Where("name = ?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user model.User) error {
	if err := infrastructure.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllStores() ([]model.Store, error) {
	var stores []model.Store
	if err := infrastructure.DB.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStoreMenues(storeID int) ([]model.Menue, error) {
	var menues []model.Menue
	if err := infrastructure.DB.Where("store_id = ?", storeID).Find(&menues).Error; err != nil {
		return nil, err
	}
	return menues, nil
}

func GetMenueAvailabil
