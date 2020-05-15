package services

import (
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repositories"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
)

// UserService -> the propose of user service is handling business logic application
type UserService struct {
	UserRepository repositories.UserRepository
}

// UService -> user service instance
func UService() UserService {
	return UserService{
		UserRepository: repositories.URepository(),
	}
}

// GetUsers -> get users service logic
func (s *UserService) GetUsers() []repositories.GetUser {
	users := s.UserRepository.GetUsers()
	return users
}

// GetUser -> get user service logic
func (s *UserService) GetUser(id int64) repositories.GetUser {
	user := s.UserRepository.GetUser(id)

	if (user == repositories.GetUser{}) {
		exception.Empty(
			flag.GetUserNotFound.Message,
			flag.GetUserNotFound.Error.Message,
			flag.GetUserNotFound.Error.Flag,
		)
	}

	return user
}

// CreateUser -> create user service logic
func (s *UserService) CreateUser(user entity.User) repositories.GetUser {
	userExist := s.UserRepository.UserExist(
		repositories.UserExistParams{Email: user.Email},
	)

	if (userExist != entity.User{}) {
		exception.BadRequest(
			flag.CreateUserAlreadyExist.Error.Message,
			flag.CreateUserAlreadyExist.Error.Flag,
		)
	}

	data := s.UserRepository.CreateUser(user)
	return data
}

// UpdateUser -> update user service logic
func (s *UserService) UpdateUser(id uint, user entity.User) repositories.GetUser {
	userExist := s.UserRepository.UserExist(repositories.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.BadRequest(
			flag.UpdateUserNotExist.Error.Message,
			flag.UpdateUserNotExist.Error.Flag,
		)
	}

	data := s.UserRepository.UpdateUser(id, user)
	return data
}

// DeleteUser -> delete user service logic
func (s *UserService) DeleteUser(id uint) repositories.GetUser {
	userExist := s.UserRepository.UserExist(repositories.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.BadRequest(
			flag.DeleteUserNotExist.Error.Message,
			flag.DeleteUserNotExist.Error.Flag,
		)
	}

	data := s.UserRepository.DeleteUser(id)
	return data
}
