package repository

import "bentol/domain/model"

type UserRepository interface {
	FindByNameAndPassword(name, password string) (*model.User, error)
}

type StoreRepository interface {
	GetAllStores() ([]model.Store, error)
}

type MenueRepository interface {
	GetMenuesByStoreID(storeID uint) ([]model.Menue, error)
	GetMenueByID(id uint) (*model.Menue, error)
}

type UserReservationRepository interface {
	CreateReservation(reservation *model.UserReservation) error
	DeleteReservationByID(id uint) error
}
