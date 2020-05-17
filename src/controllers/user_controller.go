package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/services"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
	"github.com/saefullohmaslul/golang-example/src/validations/schemas"
)

// UserController -> the propose of user controller
// is handling request from client and
// forward it to specific service
type UserController struct {
	Service services.UserService
}

// UController -> user controller instance
func UController() UserController {
	return UserController{
		Service: services.UService(),
	}
}

// GetUsers -> get users routes
// GET /users
func (u *UserController) GetUsers(c *gin.Context) {
	users := u.Service.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.GetUsersSuccess.Message,
		"data":    users,
		"error":   nil,
	})
}

// GetUser -> get user by id routes
// GET /user/:id
func (u *UserController) GetUser(c *gin.Context) {
	param := schemas.UserID{}
	_ = c.ShouldBindUri(&param)

	user := u.Service.GetUser(int64(param.ID))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.GetUserSuccess.Message,
		"data":    user,
		"error":   nil,
	})
}

// CreateUser -> create user routes
// POST /user
func (u *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	data := u.Service.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.CreateUserSuccess.Message,
		"data":    data,
		"error":   nil,
	})
}

// UpdateUser -> update user routes by id
// PATCH /user/:id
func (u *UserController) UpdateUser(c *gin.Context) {
	user := entity.User{}
	param := schemas.UserID{}
	_ = c.ShouldBindUri(&param)
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	data := u.Service.UpdateUser(param.ID, user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.UpdateUserSuccess.Message,
		"data":    data,
		"error":   nil,
	})
}

// DeleteUser -> delete user routes by id
// DELETE /user/:id
func (u *UserController) DeleteUser(c *gin.Context) {
	param := schemas.UserID{}
	_ = c.ShouldBindUri(&param)

	data := u.Service.DeleteUser(param.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flag.DeleteUserSuccess.Message,
		"data":    data,
		"error":   nil,
	})
}
