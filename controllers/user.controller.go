package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/global/types"
	"github.com/saefullohmaslul/Golang-Example/utils"
)

// UserController is controller for user module
type UserController struct {
}

// GetName will retrive name to response body
func (u UserController) GetName(c echo.Context) error {
	name := types.GetNameResponse{
		Result: "Saefulloh Maslul",
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success get name",
		},
	}

	return c.JSON(http.StatusOK, name)
}

// GetBiodata will retrive name and address to response body
func (u UserController) GetBiodata(c echo.Context) error {
	biodata := types.GetBiodataResponse{
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success to get biodata",
		},
		Result: types.GetBiodataResult{
			Name:    "Saefulloh Maslul",
			Address: "Tegal",
		},
	}

	return c.JSON(http.StatusOK, biodata)
}
