package controllers

import (
	"net/http"
	"restapi/src/models"
	"restapi/src/services"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type AccountController interface {
	CheckSaldo(echo.Context) error
}

type AccountControllerImpl struct {
	service *services.Service
}

func NewAccountController(ioc di.Container) AccountController {
	return &AccountControllerImpl{
		service: ioc.Get("service").(*services.Service),
	}
}

func (ctl *AccountControllerImpl) CheckSaldo(c echo.Context) error {
	var (
		res           models.GenericRes
		accountNumber int64
	)

	if accountNumber, res.Error = strconv.ParseInt(c.Param("account_number"), 10, 64); res.Error != nil {
		res.Code = http.StatusBadRequest
		res.Message = res.Error.(error).Error()
		return c.JSON(res.Code, res)
	}

	if res = ctl.service.Account.CheckSaldo(accountNumber); res.Error != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(res.Code, res)
}
