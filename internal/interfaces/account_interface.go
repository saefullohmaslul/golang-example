package interfaces

import (
	"restapi/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountRest interface {
	CheckBalance(echo.Context) error
	Transfer(echo.Context) error
}

type AccountService interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	Transfer(*models.TransferBalance) error
}

type AccountRoute interface {
	Setup()
}

type AccountRepository interface {
	Repository

	UseTransaction(tx *gorm.DB) AccountRepository
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	GetAccountByPks([]*int64) ([]models.Account, error)
	CheckInsufficientBalance(*int64, *int64) (models.Account, error)
	UpdateBalance(params *models.UpdateBalance) (err error)
}
