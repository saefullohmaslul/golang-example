package controllers

import (
	"restapi/src/services"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewController),
)

type Controller interface {
	CheckBalance(echo.Context) error
	Transfer(echo.Context) error
}

type ControllerImpl struct {
	service services.Service
}

func NewController(service services.Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}
