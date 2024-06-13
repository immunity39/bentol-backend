package infrastructure

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigList struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
}

var Config ConfigList
var DB *gorm.DB

func loadConfig() {
	// Get the current directory
	//Config = ConfigList{
	//    dbUser:     os.Getenv("dbUser")
	//    dbPassword: os.Getenv("dbPassword")
	//    dbHost:     os.Getenv("dbHost")
	//    dbPort:     os.Getenv("dbPort")
	//    dbName:     os.Getenv("dbName")
	//}
}

func InitDB() {
	// loadConfig()

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//    Config.dbUser, Config.dbPassword, Config.dbHost, Config.dbPort, Config.dbName)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful")
    DB.AutoMigrate(&model.User{}, &model.Store{}, &model.Menue{}, &model.UserReservation{}, &model.StoreSchedule{})
}
