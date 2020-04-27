package app

import (
	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/routes"
)

// Application struct
type Application struct {
}

// CreateApp method
func (a Application) CreateApp(e *echo.Echo) {
	configureAPIEndpoint(e)
}

func configureAPIEndpoint(e *echo.Echo) {
	g := e.Group("/user")
	routes.Router(g)
}
