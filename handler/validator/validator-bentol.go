package validator

import (
	"backend/domain/model"
	"errors"
	"net/mail"
)

func ValidateUser(user model.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if _, err := mail.ParseAddress(user.Mail); err != nil {
		return errors.New("invalid email address")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}

