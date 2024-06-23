package config

import (
	"bentol/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.Menue{},
		&models.UserReservation{},
		&models.StoreVendor{},
		&models.StoreBasicReservationPolicy{},
		&models.StoreSpecificReservationPolicy{},
		&models.StoreSchedule{},
	)

	DB = database
}

func GetDB() *gorm.DB {
	if DB == nil {
		ConnectDatabase()
	}
	return DB
}
