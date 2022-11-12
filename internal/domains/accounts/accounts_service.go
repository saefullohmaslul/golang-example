package accounts

import (
	"net/http"
	"restapi/internal/interfaces"
	"restapi/internal/models"
	"restapi/internal/repositories"

	"github.com/labstack/echo/v4"
)

type AccountServiceImpl struct {
	repository   repositories.Repository
	emailService interfaces.EmailService
}

func NewAccountService(repository repositories.Repository, emailService interfaces.EmailService) interfaces.AccountService {
	return &AccountServiceImpl{
		repository:   repository,
		emailService: emailService,
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

		if err != nil {
			return err
		}

		err = s.emailService.SendNotificationTransfer()

		return err
	})
}
