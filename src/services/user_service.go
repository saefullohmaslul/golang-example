package services

import (
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repository"
)

var (
	userRepository repository.UserRepository = repository.UserRepository{}
)

// UserService -> the propose of user service
// is handling business logic application
type UserService struct {
}

// GetUsers -> get users service logic
func (u *UserService) GetUsers() []repository.GetUser {
	users := userRepository.GetUsers()
	return users
}

// GetUser -> get user service logic
func (u *UserService) GetUser(id int64) repository.GetUser {
	user := userRepository.GetUser(id)

	if (user == repository.GetUser{}) {
		exception.Empty("User not found", "User with this ID not enough", "USER_NOT_FOUND")
	}

	return user
}

// CreateUser -> create user service logic
func (u *UserService) CreateUser(user entity.User) repository.GetUser {
	userExist := userRepository.UserExist(
		repository.UserExistParams{Email: user.Email},
	)

	if (userExist != entity.User{}) {
		exception.BadRequest("User with this email already exist", "USER_ALREADY_EXIST")
	}

	data := userRepository.CreateUser(user)
	return data
}

// UpdateUser -> update user service logic
func (u *UserService) UpdateUser(id uint, user entity.User) repository.GetUser {
	userExist := userRepository.UserExist(repository.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.BadRequest("User with this id not exist", "USER_NOT_FOUND")
	}

	data := userRepository.UpdateUser(id, user)
	return data
}

// DeleteUser -> delete user service logic
func (u *UserService) DeleteUser(id uint) repository.GetUser {
	userExist := userRepository.UserExist(repository.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.BadRequest("User with this id not exist", "USER_NOT_FOUND")
	}

	data := userRepository.DeleteUser(id)
	return data
}
