package services

import (
	"bentol/config"
	"bentol/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(name, password, mail string) (models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{Name: name, Mail: mail, Password: string(hashPassword)}
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func LoginUser(mail, password string) (models.User, error) {
	var user models.User

	if err := config.DB.Where("mail = ?", mail).First(&user).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, errors.New("incorrect password")
	}
	return user, nil
}

func GenerateJWT(userID uint) (string, error) {
	times := time.Now().Add(1 * time.Hour)
	claims := &UserClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: times.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, err
}
