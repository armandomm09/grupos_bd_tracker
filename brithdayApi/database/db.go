package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		log.Fatalf("Failed to connect to database, not in .env")
	}

	db, err := gorm.Open(mysql.Open(dburl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db
}