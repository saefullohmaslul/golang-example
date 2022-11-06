package repositories

import (
	"restapi/internal/lib"
	"restapi/internal/models"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(
	fx.Provide(NewReposiory),
)

type Repository interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	GetAccountByPks([]*int64) ([]models.Account, error)
	CheckInsufficientBalance(*int64, *int64) (models.Account, error)
	WithTransaction(f func(r Repository) error) error
	UpdateBalance(params *models.UpdateBalance) (err error)
}

type RepositoryImpl struct {
	lib.Database
}

func NewReposiory(db lib.Database) Repository {
	return &RepositoryImpl{
		Database: db,
	}
}

func (r *RepositoryImpl) WithTransaction(f func(r Repository) error) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		repo := *r
		repo.Database = lib.Database{DB: tx}
		return f(&repo)
	})
}
