package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/utils"
)

// GetNameResponse handler
type GetNameResponse struct {
	utils.Response
	Result string `json:"result"`
}

// GetName endpoint
func GetName(c echo.Context) error {
	name := GetNameResponse{
		Result: "Saefulloh Maslul",
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success get name",
		},
	}

	return c.JSON(http.StatusOK, name)
}
