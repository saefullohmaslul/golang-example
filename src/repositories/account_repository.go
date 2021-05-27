package repositories

import (
	"restapi/src/models"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CheckSaldo(int64) (models.CheckSaldoAccount, error)
}

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountReposiory(ioc di.Container) AccountRepository {
	database := getDatabase(ioc)
	return &AccountRepositoryImpl{
		db: database,
	}
}

func (r *AccountRepositoryImpl) CheckSaldo(accountNumber int64) (data models.CheckSaldoAccount, err error) {
	err = r.db.Raw(`
		SELECT 
		account_number, 
		customers.name AS customer_name, 
		balance 
		FROM accounts
		LEFT JOIN customers ON customers.customer_number = accounts.customer_number
		WHERE account_number = ?;
	`, accountNumber).Scan(&data).Error
	return
}
