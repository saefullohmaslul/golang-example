package services

import (
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/repository"
)

// UserService -> the propose of user service
// is handling business logic application
type UserService struct {
	UserRepository repository.UserRepository
}

// UService service
func UService() UserService {
	return UserService{
		UserRepository: repository.URepository(),
	}
}

// GetUsers -> get users service logic
func (s *UserService) GetUsers() []repository.GetUser {
	users := s.UserRepository.GetUsers()
	return users
}

// GetUser -> get user service logic
func (s *UserService) GetUser(id int64) repository.GetUser {
	user := s.UserRepository.GetUser(id)

	if (user == repository.GetUser{}) {
		exception.Empty("User not found", "User with this ID not enough", "USER_NOT_FOUND")
	}

	return user
}

// CreateUser -> create user service logic
func (s *UserService) CreateUser(user entity.User) repository.GetUser {
	userExist := s.UserRepository.UserExist(
		repository.UserExistParams{Email: user.Email},
	)

	if (userExist != entity.User{}) {
		exception.BadRequest("User with this email already exist", "USER_ALREADY_EXIST")
	}

	data := s.UserRepository.CreateUser(user)
	return data
}

// UpdateUser -> update user service logic
func (s *UserService) UpdateUser(id uint, user entity.User) repository.GetUser {
	userExist := s.UserRepository.UserExist(repository.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.BadRequest("User with this id not exist", "USER_NOT_FOUND")
	}

	data := s.UserRepository.UpdateUser(id, user)
	return data
}

// DeleteUser -> delete user service logic
func (s *UserService) DeleteUser(id uint) repository.GetUser {
	userExist := s.UserRepository.UserExist(repository.UserExistParams{ID: id})
	if (userExist == entity.User{}) {
		exception.BadRequest("User with this id not exist", "USER_NOT_FOUND")
	}

	data := s.UserRepository.DeleteUser(id)
	return data
}
