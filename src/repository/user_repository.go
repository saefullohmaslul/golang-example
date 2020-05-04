package repository

import (
	"github.com/jinzhu/gorm"
	db "github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
)

// UserRepository layer
type UserRepository struct {
}

func query() *gorm.DB {
	return db.GetDB().Table("users")
}

// GetUser struct
type GetUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
}

// GetUsers repository
func (u *UserRepository) GetUsers() []GetUser {
	users := []GetUser{}
	query().Select("name, email, address, age").Find(&users)
	return users
}

// GetUser repository
func (u *UserRepository) GetUser(id int64) GetUser {
	user := GetUser{}
	query().Select("name, email, address, age").Where("id = ?", id).First(&user)
	return user
}

// UserExist to check if user already exist
func (u *UserRepository) UserExist(email string) entity.User {
	user := entity.User{}
	query().Select("email").Where(&entity.User{Email: email}).First(&user)
	return user
}

// CreateUser to insert user in DB
func (u *UserRepository) CreateUser(user entity.User) GetUser {
	query().Create(&user)
	userCreated := GetUser{}
	query().Select("name, email, address, age").Where("id = ?", user.ID).First(&userCreated)
	return userCreated
}
