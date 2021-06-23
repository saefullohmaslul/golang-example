package services

import (
	"errors"
	"net/http"
	"restapi/src/constants"
	"restapi/src/models"
	"restapi/src/repositories"
	"restapi/src/utils/fixtures"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountService(t *testing.T) {
	assert := assert.New(t)

	t.Run("It should get data from ioc", func(t *testing.T) {
		builder, _ := di.NewBuilder()
		builder.Add(di.Def{
			Name: "repository",
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepositoryMock(builder.Build()), nil
			},
		})

		accountService := NewAccountService(builder.Build())
		assert.NotNil(t, accountService)
	})
}

func TestCheckBalance(t *testing.T) {
	t.Run("It should return balance account", func(t *testing.T) {
		var (
			accountNumber int64 = 1
		)

		mocks := repositories.AccountRepositoryMock{
			CheckBalanceData: fixtures.CheckBalanceAccount,
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		data, _ := accountService.CheckBalance(&accountNumber)

		assert.Equal(t, data, fixtures.CheckBalanceAccount)
	})

	t.Run("It should return account not found", func(t *testing.T) {
		var (
			accountNumber int64 = 1
		)

		mocks := repositories.AccountRepositoryMock{
			CheckBalanceData: models.CheckBalanceAccount{},
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		data, _ := accountService.CheckBalance(&accountNumber)

		assert.Equal(t, data, models.CheckBalanceAccount{})
	})
}

func TestTransfer(t *testing.T) {
	t.Run("It should transfer to another account", func(t *testing.T) {
		bodies := fixtures.TransferBalance

		mocks := repositories.AccountRepositoryMock{
			GetAccountByPksData: append(
				[]models.Account{}, fixtures.Account, fixtures.Account,
			),
			CheckInsufficientBalanceData: fixtures.Account,
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		err := accountService.Transfer(&bodies)
		assert.Equal(t, err, nil)
	})

	t.Run("It should error when get account by pks", func(t *testing.T) {
		bodies := fixtures.TransferBalance

		mocks := repositories.AccountRepositoryMock{
			GetAccountByPksError:         errors.New("some error"),
			CheckInsufficientBalanceData: fixtures.Account,
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		err := accountService.Transfer(&bodies)
		assert.Equal(t, err, errors.New("some error"))
	})

	t.Run("It should account not found", func(t *testing.T) {
		bodies := fixtures.TransferBalance

		mocks := repositories.AccountRepositoryMock{
			GetAccountByPksData: append(
				[]models.Account{}, fixtures.Account,
			),
			CheckInsufficientBalanceData: fixtures.Account,
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		err := accountService.Transfer(&bodies)
		assert.Equal(t, err, echo.NewHTTPError(http.StatusNotFound, constants.ACCOUNT_NOT_FOUND))
	})

	t.Run("It should error when check insufficient balance", func(t *testing.T) {
		bodies := fixtures.TransferBalance

		mocks := repositories.AccountRepositoryMock{
			GetAccountByPksData: append(
				[]models.Account{}, fixtures.Account, fixtures.Account,
			),
			CheckInsufficientBalanceError: errors.New("some error"),
			CheckInsufficientBalanceData:  fixtures.Account,
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		err := accountService.Transfer(&bodies)
		assert.Equal(t, err, errors.New("some error"))
	})

	t.Run("It should insufficient balance", func(t *testing.T) {
		bodies := fixtures.TransferBalance

		mocks := repositories.AccountRepositoryMock{
			GetAccountByPksData: append(
				[]models.Account{}, fixtures.Account, fixtures.Account,
			),
			CheckInsufficientBalanceData: models.Account{},
		}

		repository := &repositories.Repository{
			Account: repositories.NewAccountReposioryMock(mocks),
		}

		accountService := &AccountServiceImpl{
			repository: repository,
		}

		err := accountService.Transfer(&bodies)
		assert.Equal(t, err, echo.NewHTTPError(http.StatusBadRequest, constants.INSUFFICIENT_BALANCE))
	})
}
