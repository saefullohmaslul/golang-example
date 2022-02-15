package controllers

import (
	"net/http"
	"restapi/src/constants"
	"restapi/src/models"
	"restapi/src/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountController interface {
	CheckBalance(echo.Context) error
	Transfer(echo.Context) error
}

type AccountControllerImpl struct {
	service services.Service
}

func NewAccountController(service services.Service) AccountController {
	return &AccountControllerImpl{
		service: service,
	}
}

func (ctl *AccountControllerImpl) CheckBalance(c echo.Context) error {
	var (
		accountNumber int64
		err           error
		data          models.CheckBalanceAccount
	)

	if accountNumber, err = strconv.ParseInt(c.Param("account_number"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if data, err = ctl.service.CheckBalance(&accountNumber); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: constants.SUCCESS_CHECK_BALANCE,
		Data:    data,
	})
}

func (ctl *AccountControllerImpl) Transfer(c echo.Context) error {
	var (
		err error
	)

	bodies := new(models.TransferBalance)

	if bodies.FromAccountNumber, err = strconv.ParseInt(c.Param("from_account_number"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Bind(bodies); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(bodies); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctl.service.Transfer(bodies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
