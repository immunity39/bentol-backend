package repository

import "bentol/domain/model"

type UserRepository interface {
	FindByNameAndPassword(name, password string) (*model.User, error)
}
