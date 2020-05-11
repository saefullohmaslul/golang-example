package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/golang-example/package/kafka"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/services"
	"github.com/saefullohmaslul/golang-example/src/validation"
)

var (
	userService services.UserService = services.UserService{}
)

// UserController -> the propose of user controller
// is handling request from client and
// forward it to specific service
type UserController struct {
}

// GetUsers -> get users routes
// GET /users
func (u *UserController) GetUsers(c *gin.Context) {
	users := userService.GetUsers()

	topic := "test_topic"
	message := "Aku Bulat"
	kafka.PublishTopic(topic, message)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get all users",
		"result":  users,
	})
}

// GetUser -> get user by id routes
// GET /user/:id
func (u *UserController) GetUser(c *gin.Context) {
	param := validation.GetUserParamSchema{}
	_ = c.ShouldBindUri(&param)

	user := userService.GetUser(int64(param.ID))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get user",
		"result":  user,
	})
}

// CreateUser -> create user routes
// POST /user
func (u *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	data := userService.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success create user",
		"result":  data,
	})
}

// UpdateUser -> update user routes by id
// PATCH /user/:id
func (u *UserController) UpdateUser(c *gin.Context) {
	user := entity.User{}
	param := validation.GetUserParamSchema{}
	_ = c.ShouldBindUri(&param)
	_ = c.ShouldBindBodyWith(&user, binding.JSON)

	data := userService.UpdateUser(param.ID, user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success update user",
		"result":  data,
	})
}

// DeleteUser -> delete user routes by id
// DELETE /user/:id
func (u *UserController) DeleteUser(c *gin.Context) {
	param := validation.GetUserParamSchema{}
	_ = c.ShouldBindUri(&param)

	data := userService.DeleteUser(param.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success delete user",
		"result":  data,
	})
}
