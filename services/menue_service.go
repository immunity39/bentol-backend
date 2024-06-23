package services

import (
    "bentol/config"
    "bentol/models"
)

func AddMenue(storeID uint, name string, price uint, description string) (models.Menue, error) {
    menue := models.Menue{
        StoreID:     storeID,
        Name:        name,
        Price:       price,
        Description: description,
    }
    if err := config.DB.Create(&menue).Error; err != nil {
        return models.Menue{}, err
    }
    return menue, nil
}

func UpdateMenue(id uint, name string, price uint, description string) (models.Menue, error) {
    var menue models.Menue
    if err := config.DB.First(&menue, id).Error; err != nil {
        return models.Menue{}, err
    }
    menue.Name = name
    menue.Price = price
    menue.Description = description

    if err := config.DB.Save(&menue).Error; err != nil {
        return models.Menue{}, err
    }
    return menue, nil
}
