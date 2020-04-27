package main

import (
	"github.com/labstack/echo/v4"
	"github.com/saefullohmaslul/Golang-Example/app"
)

// main project
func main() {
	e := echo.New()

	app := new(app.Application)
	app.CreateApp(e)

	e.Logger.Fatal(e.Start(":8080"))
}
