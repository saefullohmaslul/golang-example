package validations

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
	"github.com/saefullohmaslul/golang-example/src/utils/flag"
	"github.com/saefullohmaslul/golang-example/src/validations/schemas"
)

// CreateUser -> validations to create user
func CreateUser(c *gin.Context) {
	var user schemas.CreateUser
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		exception.BadRequest(err.Error(), flag.CreateUserInvalidBody.Error.Flag)
	}

	userValidate := &schemas.CreateUser{
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
	}
	Validate(userValidate)
}

// GetUser -> validations to get user by id
func GetUser(c *gin.Context) {
	param := schemas.UserID{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest(
			flag.GetUserInvalidParamID.Error.Message,
			flag.GetUserInvalidParamID.Error.Flag,
		)
	}
}

// UpdateUser -> validations to update user by id with body
func UpdateUser(c *gin.Context) {
	param := schemas.UserID{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest(
			flag.UpdateUserInvlidParamURI.Error.Message,
			flag.UpdateUserInvlidParamURI.Error.Flag,
		)
	}

	var user schemas.UpdateUser
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		exception.BadRequest(err.Error(), flag.UpdateUserInvalidBody.Error.Flag)
	}

	userValidate := &schemas.UpdateUser{
		Name:    user.Name,
		Address: user.Address,
		Age:     user.Age,
		Email:   user.Email,
	}
	Validate(userValidate)
}

// DeleteUser -> validations to delete user by id
func DeleteUser(c *gin.Context) {
	param := schemas.UserID{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest(
			flag.DeleteUserInvalidParamURI.Error.Message,
			flag.DeleteUserInvalidParamURI.Error.Flag,
		)
	}
}
