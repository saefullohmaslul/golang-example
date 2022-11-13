package accounts

import (
	"net/http"
	"restapi/internal/interfaces"
	"restapi/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountRestImpl struct {
	service interfaces.AccountService
}

func NewAccountRest(service interfaces.AccountService) interfaces.AccountRest {
	return &AccountRestImpl{
		service: service,
	}
}

func (ctl *AccountRestImpl) CheckBalance(c echo.Context) error {
	var (
		accountNumber int64
		err           error
		data          models.CheckBalanceAccount
		ctx           = c.Request().Context()
	)

	if accountNumber, err = strconv.ParseInt(c.Param("account_number"), 10, 64); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if data, err = ctl.service.CheckBalance(ctx, accountNumber); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.GenericRes{
		Code:    http.StatusOK,
		Message: SUCCESS_CHECK_BALANCE,
		Data:    data,
	})
}

func (ctl *AccountRestImpl) Transfer(c echo.Context) error {
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
