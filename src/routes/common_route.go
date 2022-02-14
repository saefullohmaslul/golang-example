package routes

import (
	"net/http"
	"restapi/src/lib"
	"restapi/src/models"

	"github.com/labstack/echo/v4"
)

type CommonRouter struct {
	echoHandler lib.EchoHandler
}

func (r *CommonRouter) Setup() {
	router := r.echoHandler.Echo.Group("")
	{
		router.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusOK,
				Message: "OK.",
			})
		})
		router.Any("*", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusNotFound,
				Message: "Route not found.",
			})
		})
	}
}
