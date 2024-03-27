package database

import (
	"fmt"
	"goGram/models"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	debugMode bool
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	debugModeStr := os.Getenv("DEBUG_MODE")
	debugMode, _ = strconv.ParseBool(debugModeStr)
}

func StartDB() {
	loadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if debugMode {
		runMigrations()
	}
}

func GetDB() *gorm.DB {
	return db.Debug()
}

func runMigrations() {
	err := db.AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}
}
