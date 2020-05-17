package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/saefullohmaslul/golang-example/src/middlewares/exception"
)

// Validate -> function to validate request
func Validate(schema interface{}, errors []map[string]interface{}) {
	/**
	 * create validator instance
	 */
	validate := validator.New()

	if err := validate.Struct(schema); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errors = append(errors, map[string]interface{}{
				"message": fmt.Sprint(err), "flag": "INVALID_BODY"},
			)
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, map[string]interface{}{
				"message": fmt.Sprint(err), "flag": "INVALID_BODY"},
			)
		}
		exception.BadRequest("Validation error", errors)
	}
	if errors != nil {
		exception.BadRequest("Validation error", errors)
	}
}
