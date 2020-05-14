package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
)

// CreateUser is create user tabel for migration
func CreateUser(conn *gorm.DB) {
	conn.CreateTable(&entity.User{})
}
