package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
)

// Migrate -> execution migration based on the table
func Migrate(conn *gorm.DB) {
	conn.AutoMigrate(entity.User{})
}
