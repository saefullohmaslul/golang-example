package services

import (
	"net/http"
	"restapi/src/constants"
	"restapi/src/models"
	"restapi/src/repositories"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type AccountService interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	Transfer(*models.TransferBalance) error
}

type AccountServiceImpl struct {
	repository *repositories.Repository
}

func NewAccountService(ioc di.Container) AccountService {
	return &AccountServiceImpl{
		repository: ioc.Get("repository").(*repositories.Repository),
	}
}

func (s *AccountServiceImpl) CheckBalance(accountNumber *int64) (data models.CheckBalanceAccount, err error) {
	data, err = s.repository.Account.CheckBalance(accountNumber)

	if data.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusNotFound, constants.ACCOUNT_NOT_FOUND)
		return
	}
	return
}

func (s *AccountServiceImpl) Transfer(bodies *models.TransferBalance) (err error) {
	var (
		accounts []models.Account
		account  models.Account
	)

	if accounts, err = s.repository.Account.GetAccountByPks(
		[]*int64{&bodies.FromAccountNumber, &bodies.ToAccountNumber},
	); err != nil {
		return
	}

	if len(accounts) < 2 {
		err = echo.NewHTTPError(http.StatusNotFound, constants.ACCOUNT_NOT_FOUND)
		return
	}

	if account, err = s.repository.Account.CheckInsufficientBalance(&bodies.FromAccountNumber, &bodies.Amount); err != nil {
		return
	}

	if account.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, constants.INSUFFICIENT_BALANCE)
		return
	}

	err = s.repository.Account.TransferBalance(bodies)
	return
}
