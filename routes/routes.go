package routes

import (
	"github.com/saefullohmaslul/Golang-Example/controllers"

	"github.com/labstack/echo/v4"
)

// Router middleware
func Router(g *echo.Group) {
	g.GET("/name", controllers.GetName)
	g.GET("/biodata", controllers.GetBiodata)
}
