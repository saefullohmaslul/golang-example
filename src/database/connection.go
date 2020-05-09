package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
)

var conn *gorm.DB
var err error

// Connection -> create connection with credentials
func Connection() {
	authDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	conn, err = gorm.Open("postgres", authDB)
	if err != nil {
		logging.Error("DB", err)
	}
}

// AppConnection -> method to create connection for application
func AppConnection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// TestConnection -> method to create connection for application testing
func TestConnection() {
	if err := godotenv.Load("../.env"); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// GetDB -> method to get connection instance
func GetDB() *gorm.DB {
	return conn
}

// DropAllTable -> method to drop all database table (using this only for testing)
func DropAllTable() {
	conn.DropTable(&entity.User{})
}
