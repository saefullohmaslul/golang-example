package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/database"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/global/types"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/validation"
)

// UserController is controller for user module
type UserController struct {
}

// GetUsers will retrive all user
func (u UserController) GetUsers(c *gin.Context) {
	users := []entity.User{}
	database.GetDB().Select("name, email, address, age").Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get all users",
		"result":  users,
	})
}

type getUser struct {
	ID int `uri:"id" binding:"required"`
}

// GetUser will retrive user
func (u UserController) GetUser(c *gin.Context) {
	param := getUser{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Param must be of type integer, required", "INVALID_BODY")
	}

	user := entity.User{}
	database.GetDB().Select("name, email, address, age").First(&user, param.ID)

	if (user == entity.User{}) {
		exception.Empty("User not found", "User with this ID not enough", "USER_NOT_FOUND")
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get user",
		"result":  user,
	})
}

// CreateUser will add user into database
func (u UserController) CreateUser(c *gin.Context) {
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

	isUsed := entity.User{}
	database.GetDB().Where(&entity.User{Email: user.Email}).First(&isUsed)
	if (isUsed != entity.User{}) {
		exception.BadRequest("User already exist", "USER_ALREADY_EXIST")
	}

	if dbc := database.GetDB().Create(&user); dbc.Error != nil {
		exception.InternalServerError("Can't create user", "DATABASE_ERROR")
	}

	data := types.CreateUserResult{
		ID:      user.ID,
		Name:    user.Name,
		Age:     user.Age,
		Email:   user.Email,
		Address: user.Address,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success create user",
		"result":  data,
	})
}
