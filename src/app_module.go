package main

import "github.com/labstack/echo/v4"

type Module struct {
}

func (m Module) New(e *echo.Echo) {
	s := NewService()

	s.NewRoute(e)
}
