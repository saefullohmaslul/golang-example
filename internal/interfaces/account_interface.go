package interfaces

import (
	"context"
	"restapi/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountRest interface {
	CheckBalance(echo.Context) error
	Transfer(echo.Context) error
}

type AccountService interface {
	CheckBalance(ctx context.Context, accountNumber int64) (data models.CheckBalanceAccount, err error)
	Transfer(ctx context.Context, bodies *models.TransferBalance) (err error)
}

type AccountRoute interface {
	Setup()
}

type AccountRepository interface {
	Repository

	UseTransaction(tx *gorm.DB) AccountRepository
	CheckBalance(ctx context.Context, accountNumber int64) (data models.CheckBalanceAccount, err error)
	GetAccountByPks(ctx context.Context, accountPks []int64) ([]models.Account, error)
	CheckInsufficientBalance(ctx context.Context, accountNumber, amount int64) (data models.Account, err error)
	UpdateBalance(ctx context.Context, params *models.UpdateBalance) (err error)
}
