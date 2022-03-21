package accounts

import (
	"net/http"
	"restapi/src/models"
	"restapi/src/repositories"

	"github.com/labstack/echo/v4"
)

type AccountService interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	Transfer(*models.TransferBalance) error
}

type AccountServiceImpl struct {
	repository repositories.Repository
}

func NewAccountService(repository repositories.Repository) AccountService {
	return &AccountServiceImpl{
		repository: repository,
	}
}

func (s *AccountServiceImpl) CheckBalance(accountNumber *int64) (data models.CheckBalanceAccount, err error) {
	data, err = s.repository.CheckBalance(accountNumber)

	if data.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusNotFound, ACCOUNT_NOT_FOUND)
		return
	}
	return
}

func (s *AccountServiceImpl) Transfer(bodies *models.TransferBalance) (err error) {
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
		err = echo.NewHTTPError(http.StatusNotFound, ACCOUNT_NOT_FOUND)
		return
	}

	if account, err = s.repository.CheckInsufficientBalance(&bodies.FromAccountNumber, &bodies.Amount); err != nil {
		return
	}

	if account.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, INSUFFICIENT_BALANCE)
		return
	}

	return s.TransferBalance(bodies)
}

func (s *AccountServiceImpl) TransferBalance(bodies *models.TransferBalance) (err error) {
	return s.repository.WithTransaction(func(r repositories.Repository) error {
		if err = r.UpdateBalance(&models.UpdateBalance{
			AccountNumber: bodies.ToAccountNumber,
			Amount:        bodies.Amount,
		}); err != nil {
			return err
		}

		err = r.UpdateBalance(&models.UpdateBalance{
			AccountNumber: bodies.FromAccountNumber,
			Amount:        -bodies.Amount,
		})

		return err
	})
}
