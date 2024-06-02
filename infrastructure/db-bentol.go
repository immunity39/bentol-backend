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
