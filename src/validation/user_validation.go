package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/saefullohmaslul/golang-example/src/database/entity"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
)

// CreateUserSchema -> create user schema validation
type CreateUserSchema struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Age      int64
	Address  string
}

// CreateUser validation
func CreateUser(c *gin.Context) {
	var user entity.User
	_ = c.BindJSON(&user)

	userValidate := &CreateUserSchema{
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
	}
	Validate(userValidate)
}

// GetUserParamSchema to check schema param
type GetUserParamSchema struct {
	ID int `uri:"id" binding:"required"`
}

// GetUser validation
func GetUser(c *gin.Context) {
	param := GetUserParamSchema{}
	if err := c.ShouldBindUri(&param); err != nil {
		exception.BadRequest("Param must be of type integer, required", "INVALID_BODY")
	}
}
