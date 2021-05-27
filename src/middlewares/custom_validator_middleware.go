package middlewares

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator(v *validator.Validate) *CustomValidator {
	return &CustomValidator{
		Validator: v,
	}
}

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return nil
}
