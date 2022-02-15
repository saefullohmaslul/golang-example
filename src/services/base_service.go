package services

import (
	"restapi/src/models"
	"restapi/src/repositories"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewService),
)

type Service interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	Transfer(*models.TransferBalance) error
}

type ServiceImpl struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}