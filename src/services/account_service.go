package services

import (
	"net/http"
	"restapi/src/constants"
	"restapi/src/models"
	"restapi/src/repositories"

	"github.com/sarulabs/di"
)

type AccountService interface {
	CheckSaldo(int64) models.GenericRes
}

type AccountServiceImpl struct {
	repository *repositories.Repository
}

func NewAccountService(ioc di.Container) AccountService {
	return &AccountServiceImpl{
		repository: ioc.Get("repository").(*repositories.Repository),
	}
}

func (s *AccountServiceImpl) CheckSaldo(accountNumber int64) (res models.GenericRes) {
	res.Data, res.Error = s.repository.Account.CheckSaldo(accountNumber)

	if res.Data.(models.CheckSaldoAccount).AccountNumber == 0 {
		res.Code = http.StatusNotFound
		res.Message = constants.ACCOUNT_NOT_FOUND
		return
	}

	res.Code = http.StatusOK
	res.Message = constants.SUCCESS_CHECK_BALANCE
	return
}
