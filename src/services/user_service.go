package service

import (
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repository"
)

// UserService layer
type UserService struct {
}

// GetUsers service
func (u *UserService) GetUsers() []repository.GetUser {
	userRepository := repository.UserRepository{}
	users := userRepository.GetUsers()
	return users
}

// GetUser service
func (u *UserService) GetUser(id int64) repository.GetUser {
	userRepository := repository.UserRepository{}
	user := userRepository.GetUser(id)

	if (user == repository.GetUser{}) {
		exception.Empty("User not found", "User with this ID not enough", "USER_NOT_FOUND")
	}

	return user
}

// CreateUser service
func (u *UserService) CreateUser(user entity.User) repository.GetUser {
	userRepository := repository.UserRepository{}

	userExist := userRepository.UserExist(user.Email)
	if (userExist != entity.User{}) {
		exception.BadRequest("User with this email already exist", "USER_ALREADY_EXIST")
	}

	data := userRepository.CreateUser(user)
	return data
}
