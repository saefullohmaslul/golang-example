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

// Connection conn instance
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

// AppConnection to handle connection db for app
func AppConnection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// TestConnection to handle connection db for test
func TestConnection() {
	if err := godotenv.Load("../.env"); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// GetDB connection
func GetDB() *gorm.DB {
	return conn
}

// DropAllTable for testing
func DropAllTable() {
	conn.DropTable(&entity.User{})
}
