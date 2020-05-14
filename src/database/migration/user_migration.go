package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/sirupsen/logrus"
)

// CreateUser is create user tabel for migration
func CreateUser(conn *gorm.DB) {
	conn.AutoMigrate(&entity.User{})

	logrus.Info("Success running migration")
}
