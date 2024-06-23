package services

import (
    "bentol/config"
    "bentol/models"
    "errors"
)

func RegisterUser(name, password, mail string) (models.User, error) {
    user := models.User{Name: name, Mail: mail, Password: password}
    if err := config.DB.Create(&user).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}

func LoginUser(name, password string) (models.User, error) {
    var user models.User
    if err := config.DB.Where("name = ? AND password = ?", name, password).First(&user).Error; err != nil {
        return models.User{}, errors.New("invalid credentials")
    }
    return user, nil
}
