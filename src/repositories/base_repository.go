package repositories

import (
	"restapi/src/lib"
	"restapi/src/models"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewReposiory),
)

type Repository interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	GetAccountByPks([]*int64) ([]models.Account, error)
	CheckInsufficientBalance(*int64, *int64) (models.Account, error)
	TransferBalance(*models.TransferBalance) error
}

type RepositoryImpl struct {
	lib.Database
}

func NewReposiory(db lib.Database) Repository {
	return &RepositoryImpl{
		Database: db,
	}
}
