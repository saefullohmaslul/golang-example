package main

import (
	"net/http"
	"restapi/src/controllers"
	"restapi/src/models"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Service struct {
	controller *controllers.Controller
}

func NewService(ioc di.Container) *Service {
	return &Service{
		controller: ioc.Get("controller").(*controllers.Controller),
	}
}

func (s *Service) NewRoute(e *echo.Echo) {
	g := e.Group("")

	s.HealthRoute(g)
	s.NotFoundRoute(g)

	g.GET("/account/:account_number", s.controller.Account.CheckSaldo)
}

func (*Service) HealthRoute(g *echo.Group) {
	{
		g.GET("/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusOK,
				Message: "OK.",
			})
		})
	}
}

func (*Service) NotFoundRoute(g *echo.Group) {
	{
		g.Any("*", func(c echo.Context) error {
			return c.JSON(http.StatusOK, models.GenericRes{
				Code:    http.StatusNotFound,
				Message: "Route not found.",
			})
		})
	}
}
