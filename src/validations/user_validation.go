package validations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/validations/schemas"
)

// CreateUser -> validations to create user
func CreateUser(c *gin.Context) {
	var errors []map[string]interface{}
	var user schemas.CreateUser
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	userValidate := &schemas.CreateUser{
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
	}
	Validate(userValidate, errors)
}

// GetUser -> validations to get user by id
func GetUser(c *gin.Context) {
	param := schemas.UserID{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Validation error", []map[string]interface{}{
			{"message": "Param must be of type integer, required", "flag": "INVALID_BODY"},
		})
	}
}

// UpdateUser -> validations to update user by id with body
func UpdateUser(c *gin.Context) {
	var errors []map[string]interface{}
	param := schemas.UserID{}
	if err := c.ShouldBindUri(&param); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": "Param must be of type integer, required", "flag": "INVALID_BODY"},
		)
	}

	var user schemas.UpdateUser
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	userValidate := &schemas.UpdateUser{
		Name:    user.Name,
		Address: user.Address,
		Age:     user.Age,
		Email:   user.Email,
	}
	Validate(userValidate, errors)
}

// DeleteUser -> validations to delete user by id
func DeleteUser(c *gin.Context) {
	param := schemas.UserID{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Validation error", []map[string]interface{}{
			{"message": "Param must be of type integer, required", "flag": "INVALID_BODY"},
		})
	}
}
