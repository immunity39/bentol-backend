package config

import (
	"bentol/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:password@tcp(127.0.0.1:3306)/bentol_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Menue{},
		&models.UserReservation{},
		&models.StoreReservationPolicy{},
		&models.StoreTimeSlotReservation{},
	)

	DB = database
}
