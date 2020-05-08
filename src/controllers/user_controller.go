package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/services"
	"github.com/saefullohmaslul/golang-example/src/validation"
)

var (
	userService services.UserService = services.UserService{}
)

// UserController is controller for user module
type UserController struct {
}

// GetUsers will retrieve all user
func (u *UserController) GetUsers(c *gin.Context) {
	users := userService.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get all users",
		"result":  users,
	})
}

// GetUser will retrieve user
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

// CreateUser will add user into database
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

// UpdateUser will update user by id
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

// DeleteUser will delete user by id
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
