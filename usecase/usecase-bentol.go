package usecase

import (
	"bentol/domain/model"
	"bentol/domain/repository"
)

type LoginUsecase struct {
	UserRepository repository.UserRepository
}

func NewLoginUsecase(ur repository.UserRepository) *LoginUsecase {
	return &LoginUsecase{UserRepository: ur}
}

func (lu *LoginUsecase) Login(name, password string) (*model.User, error) {
	return lu.UserRepository.FindByNameAndPassword(name, password)
}

type StoreUsecase struct {
	StoreRepository repository.StoreRepository
}

func NewStoreUsecase(sr repository.StoreRepository) *StoreUsecase {
	return &StoreUsecase{StoreRepository: sr}
}

func (su *StoreUsecase) GetAllStores() ([]model.Store, error) {
	return su.StoreRepository.GetAllStores()
}

type MenueUsecase struct {
	MenueRepository repository.MenueRepository
}

func NewMenueUsecase(mr repository.MenueRepository) *MenueUsecase {
	return &MenueUsecase{MenueRepository: mr}
}

func (mu *MenueUsecase) GetMenuesByStoreID(storeID uint) ([]model.Menue, error) {
	return mu.MenueRepository.GetMenuesByStoreID(storeID)
}

func (mu *MenueUsecase) GetMenueByID(id uint) (*model.Menue, error) {
	return mu.MenueRepository.GetMenueByID(id)
}

type UserReservationUsecase struct {
	UserReservationRepository repository.UserReservationRepository
}

func NewUserReservationUsecase(urr repository.UserReservationRepository) *UserReservationUsecase {
	return &UserReservationUsecase{UserReservationRepository: urr}
}

func (uru *UserReservationUsecase) CreateReservation(reservation *model.UserReservation) error {
	return uru.UserReservationRepository.CreateReservation(reservation)
}

func (uru *UserReservationUsecase) DeleteReservationByID(id uint) error {
	return uru.UserReservationRepository.DeleteReservationByID(id)
}
