package interfaces

import (
	"restapi/internal/models"

	"github.com/labstack/echo/v4"
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
