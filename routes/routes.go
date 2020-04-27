package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/controllers"
)

// Router middleware to handler routes
func Router(g *echo.Group) {
	controller := controllers.UserController{}
	g.GET("/name", controller.GetName)
	g.GET("/biodata", controller.GetBiodata)
}
