package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
)

// AutoMigration is auto migrate database
func AutoMigration(conn *gorm.DB) {
	conn.AutoMigrate(entity.User{})
}
