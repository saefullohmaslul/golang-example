package services

import (
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repository"
)

var (
	userRepository repository.UserRepository = repository.UserRepository{}
)

// UserService layer
type UserService struct {
}

// GetUsers service
func (u *UserService) GetUsers() []repository.GetUser {
	users := userRepository.GetUsers()
	return users
}

// GetUser service
func (u *UserService) GetUser(id int64) repository.GetUser {
	user := userRepository.GetUser(id)

	if (user == repository.GetUser{}) {
		exception.Empty("User not found", "User with this ID not enough", "USER_NOT_FOUND")
	}

	return user
}

// CreateUser service
func (u *UserService) CreateUser(user entity.User) repository.GetUser {
	userExist := userRepository.UserExist(user.Email)
	if (userExist != entity.User{}) {
		exception.BadRequest("User with this email already exist", "USER_ALREADY_EXIST")
	}

	data := userRepository.CreateUser(user)
	return data
}
