package services

import (
	"net/http"
	"restapi/src/constants"
	"restapi/src/models"

	"github.com/labstack/echo/v4"
)

func (s *ServiceImpl) CheckBalance(accountNumber *int64) (data models.CheckBalanceAccount, err error) {
	data, err = s.repository.CheckBalance(accountNumber)

	if data.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusNotFound, constants.ACCOUNT_NOT_FOUND)
		return
	}
	return
}

func (s *ServiceImpl) Transfer(bodies *models.TransferBalance) (err error) {
	var (
		accounts []models.Account
		account  models.Account
	)

	if accounts, err = s.repository.GetAccountByPks(
		[]*int64{&bodies.FromAccountNumber, &bodies.ToAccountNumber},
	); err != nil {
		return
	}

	if len(accounts) < 2 {
		err = echo.NewHTTPError(http.StatusNotFound, constants.ACCOUNT_NOT_FOUND)
		return
	}

	if account, err = s.repository.CheckInsufficientBalance(&bodies.FromAccountNumber, &bodies.Amount); err != nil {
		return
	}

	if account.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, constants.INSUFFICIENT_BALANCE)
		return
	}

	err = s.repository.TransferBalance(bodies)
	return
}
