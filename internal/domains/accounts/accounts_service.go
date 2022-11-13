package accounts

import (
	"context"
	"net/http"
	"restapi/internal/interfaces"
	"restapi/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountServiceImpl struct {
	accountRepository interfaces.AccountRepository
	emailService      interfaces.EmailService
}

func NewAccountService(accountRepository interfaces.AccountRepository, emailService interfaces.EmailService) interfaces.AccountService {
	return &AccountServiceImpl{
		accountRepository: accountRepository,
		emailService:      emailService,
	}
}

func (s *AccountServiceImpl) CheckBalance(ctx context.Context, accountNumber int64) (data models.CheckBalanceAccount, err error) {
	data, err = s.accountRepository.CheckBalance(ctx, accountNumber)

	if data.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusNotFound, ACCOUNT_NOT_FOUND)
		return
	}
	return
}

func (s *AccountServiceImpl) Transfer(ctx context.Context, params *models.TransferBalance) (err error) {
	var (
		accounts []models.Account
		account  models.Account
	)

	if accounts, err = s.accountRepository.GetAccountByPks(ctx, []int64{params.FromAccountNumber, params.ToAccountNumber}); err != nil {
		return
	}

	if len(accounts) < 2 {
		err = echo.NewHTTPError(http.StatusNotFound, ACCOUNT_NOT_FOUND)
		return
	}

	if account, err = s.accountRepository.CheckInsufficientBalance(ctx, params.FromAccountNumber, params.Amount); err != nil {
		return
	}

	if account.AccountNumber == 0 {
		err = echo.NewHTTPError(http.StatusBadRequest, INSUFFICIENT_BALANCE)
		return
	}

	return s.transferBalance(ctx, params)
}

func (s *AccountServiceImpl) transferBalance(ctx context.Context, params *models.TransferBalance) (err error) {
	return s.accountRepository.Transaction(func(tx *gorm.DB) error {
		if err = s.accountRepository.UseTransaction(tx).UpdateBalance(ctx, &models.UpdateBalance{
			AccountNumber: params.ToAccountNumber,
			Amount:        params.Amount,
		}); err != nil {
			return err
		}

		err = s.accountRepository.UseTransaction(tx).UpdateBalance(ctx, &models.UpdateBalance{
			AccountNumber: params.FromAccountNumber,
			Amount:        -params.Amount,
		})

		if err != nil {
			return err
		}

		err = s.emailService.SendNotificationTransfer()
		return err
	})

}
