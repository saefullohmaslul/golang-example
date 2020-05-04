package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	service "github.com/saefullohmaslul/golang-example/src/services"
	"github.com/saefullohmaslul/golang-example/src/validation"
)

// UserController is controller for user module
type UserController struct {
}

// GetUsers will retrieve all user
func (u UserController) GetUsers(c *gin.Context) {
	users := new(service.UserService).GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get all users",
		"result":  users,
	})
}

type getUser struct {
	ID int `uri:"id" binding:"required"`
}

// GetUser will retrieve user
func (u UserController) GetUser(c *gin.Context) {
	param := getUser{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Param must be of type integer, required", "INVALID_BODY")
	}

	user := new(service.UserService).GetUser(int64(param.ID))

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get user",
		"result":  user,
	})
}

// CreateUser will add user into database
func (u UserController) CreateUser(c *gin.Context) {
	userService := service.UserService{}
	var user entity.User
	_ = c.BindJSON(&user)

	userValidate := &validation.CreateUserSchema{
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
	}
	validation.Validate(userValidate)

	data := userService.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success create user",
		"result":  data,
	})
}
