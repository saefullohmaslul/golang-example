package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

var db *gorm.DB
var err error

// Connection database instance
func Connection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}

	authDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	db, err = gorm.Open("postgres", authDB)
	if err != nil {
		logging.Error("DB", err)
	}
}

// GetDB connection
func GetDB() *gorm.DB {
	return db
}
