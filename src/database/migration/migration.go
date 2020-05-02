package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/Golang-Example/src/database/entity"
)

// Migrate database table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
