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

// GetBiodataResult is return to get biodata format
type GetBiodataResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GetBiodataResponse is response handler
type GetBiodataResponse struct {
	utils.Response
	Result GetBiodataResult `json:"result"`
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

// GetBiodata is handler to get user bio
func GetBiodata(c echo.Context) error {
	biodata := GetBiodataResponse{
		Response: utils.Response{
			Status:  http.StatusOK,
			Message: "Success to get biodata",
		},
		Result: GetBiodataResult{
			Name:    "Saefulloh Maslul",
			Address: "Tegal",
		},
	}

	return c.JSON(http.StatusOK, biodata)
}
