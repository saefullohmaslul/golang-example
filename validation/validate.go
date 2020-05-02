package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/saefullohmaslul/Golang-Example/middlewares/exception"
)

// Validate -> function to validate request
func Validate(schema interface{}) {
	validate := validator.New()

	if err := validate.Struct(schema); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			exception.BadRequest(fmt.Sprint(err), "INVALID_BODY")
		}

		for _, err := range err.(validator.ValidationErrors) {
			exception.BadRequest(fmt.Sprint(err), "INVALID_BODY")
		}
	}
}
