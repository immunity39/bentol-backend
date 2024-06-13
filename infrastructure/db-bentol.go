package infrastructure

import (
	"bentol/domain/model"
	"bentol/domain/repository"
)

type UserRepositoryImpl struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByNameAndPassword(name, password string) (*model.User, error) {
	var user model.User
	result := DB.Raw("SELECT * FROM `User` WHERE name = ? AND password = ?", name, password).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

type StoreRepositoryImpl struct{}

func NewStoreRepository() repository.StoreRepository {
	return &StoreRepositoryImpl{}
}

func (r *StoreRepositoryImpl) GetAllStores() ([]model.Store, error) {
	var stores []model.Store
	result := DB.Find(&stores)
	if result.Error != nil {
		return nil, result.Error
	}
	return stores, nil
}

type MenueRepositoryImpl struct{}

func NewMenueRepository() repository.MenueRepository {
	return &MenueRepositoryImpl{}
}

func (r *MenueRepositoryImpl) GetMenuesByStoreID(storeID uint) ([]model.Menue, error) {
	var menues []model.Menue
	result := DB.Where("store_id = ?", storeID).Find(&menues)
	if result.Error != nil {
		return nil, result.Error
	}
	return menues, nil
}

func (r *MenueRepositoryImpl) GetMenueByID(id uint) (*model.Menue, error) {
	var menue model.Menue
	result := DB.First(&menue, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &menue, nil
}

type UserReservationRepositoryImpl struct{}

func NewUserReservationRepository() repository.UserReservationRepository {
	return &UserReservationRepositoryImpl{}
}

func (r *UserReservationRepositoryImpl) CreateReservation(reservation *model.UserReservation) error {
	result := DB.Create(reservation)
	return result.Error
}

func (r *UserReservationRepositoryImpl) DeleteReservationByID(id uint) error {
	result := DB.Delete(&model.UserReservation{}, id)
	return result.Error
}
