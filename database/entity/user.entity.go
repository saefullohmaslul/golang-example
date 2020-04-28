package entity

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string
	Age      int64
	Email    string `gorm:"type:varchar(100);unique_index"`
	Address  string `gorm:"index:addr"`
	Password string `gorm:"type:varchar(255)"`
}
