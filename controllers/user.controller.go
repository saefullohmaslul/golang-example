package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/Golang-Example/database/entity"
	"github.com/saefullohmaslul/Golang-Example/global/types"
	"github.com/saefullohmaslul/Golang-Example/validation"
)

// UserController is controller for user module
type UserController struct {
}

// GetName will retrive name to response body
func (u UserController) GetName(c *gin.Context) {
	name := "Saefulloh Maslul"

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get name",
		"result":  name,
	})
}

// GetBiodata will retrive name and address to response body
func (u UserController) GetBiodata(c *gin.Context) {
	biodata := types.GetBiodataResult{
		Name:    "Saefulloh Maslul",
		Address: "Tegal",
	}

	// exception.BadRequest("error cuy", "INVALID")

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get biodata",
		"result":  biodata,
	})
}

// CreateUser will add user into database
func (u UserController) CreateUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)

	userValidate := &validation.CreateUserSchema{Name: user.Name}
	validation.Validate(userValidate)

	// database.GetDB().Create(&user)

	data := types.CreateUserResult{ID: user.ID, Name: user.Name, Age: user.Age, Email: user.Email, Address: user.Address}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get name",
		"result":  data,
	})
}
