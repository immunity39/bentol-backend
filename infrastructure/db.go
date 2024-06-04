package infrastructure

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigList struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

var Config ConfigList
var DB *gorm.DB

func loadConfig() {
	// Get the current directory
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// Construct the config file path
	configPath := filepath.Join(execDir, "infrastructure", "config.ini")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config.ini file does not exist at path: %s", configPath)
	}

	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	Config = ConfigList{
		User:     cfg.Section("db").Key("user").String(),
		Password: cfg.Section("db").Key("password").String(),
		Host:     cfg.Section("db").Key("host").String(),
		Port:     cfg.Section("db").Key("port").String(),
		DBName:   cfg.Section("db").Key("dbname").String(),
	}
}

func InitDB() {
	loadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.User, Config.Password, Config.Host, Config.Port, Config.DBName)

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful")
}
