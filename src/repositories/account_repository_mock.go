package repositories

import (
	"restapi/src/models"
)

type AccountRepositoryMock struct {
	CheckBalanceData              models.CheckBalanceAccount
	GetAccountByPksData           []models.Account
	CheckInsufficientBalanceData  models.Account
	GetAccountByPksError          error
	CheckInsufficientBalanceError error
}

func NewAccountReposioryMock(data AccountRepositoryMock) AccountRepository {
	return &data
}

func (r *AccountRepositoryMock) CheckBalance(*int64) (data models.CheckBalanceAccount, err error) {
	data = r.CheckBalanceData
	return
}

func (r *AccountRepositoryMock) GetAccountByPks([]*int64) (data []models.Account, err error) {
	data = r.GetAccountByPksData
	err = r.GetAccountByPksError
	return
}

func (r *AccountRepositoryMock) CheckInsufficientBalance(accountNumber, amount *int64) (data models.Account, err error) {
	data = r.CheckInsufficientBalanceData
	err = r.CheckInsufficientBalanceError
	return
}

func (r *AccountRepositoryMock) TransferBalance(*models.TransferBalance) (err error) {
	return
}
