package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/Golang-Example/database"
	"github.com/saefullohmaslul/Golang-Example/database/entity"
	"github.com/saefullohmaslul/Golang-Example/global/types"
	"github.com/saefullohmaslul/Golang-Example/middlewares/exception"
	"github.com/saefullohmaslul/Golang-Example/validation"
)

// UserController is controller for user module
type UserController struct {
}

// GetName will retrieve name to response body
func (u UserController) GetName(c *gin.Context) {
	name := "Saefulloh Maslul"

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get name",
		"result":  name,
	})
}

// GetBiodata will retrieve name and address to response body
func (u UserController) GetBiodata(c *gin.Context) {
	biodata := types.GetBiodataResult{
		Name:    "Saefulloh Maslul",
		Address: "Tegal",
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get biodata",
		"result":  biodata,
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
