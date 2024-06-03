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
