package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB
var err error

func Create(dsn string) (*gorm.DB, error) {
	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return connection, err
}

func Get() *gorm.DB {
	return connection
}
