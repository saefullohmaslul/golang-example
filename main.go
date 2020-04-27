package main

import (
	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/routes"
)

func main() {
	e := echo.New()

	userEndpoints(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func userEndpoints(e *echo.Echo) {
	g := e.Group("/user")
	routes.Router(g)
}
