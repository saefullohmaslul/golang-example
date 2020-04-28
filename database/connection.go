package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Connection database instance
func Connection() {
	authDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	db, err = gorm.Open("postgres", authDB)
	if err != nil {
		fmt.Println("Error connect database")
		fmt.Println(err.Error())
	}
}

// GetDB connection
func GetDB() *gorm.DB {
	return db
}
