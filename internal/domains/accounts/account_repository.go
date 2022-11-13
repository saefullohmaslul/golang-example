package accounts

import (
	"context"
	"restapi/internal/interfaces"
	"restapi/internal/lib"
	"restapi/internal/models"

	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	interfaces.RepositoryImpl
}

func NewAccountRepository(db lib.Database) interfaces.AccountRepository {
	accountRepository := &accountRepositoryImpl{}
	accountRepository.Database = db

	return accountRepository
}

func (r *accountRepositoryImpl) UseTransaction(tx *gorm.DB) interfaces.AccountRepository {
	database := lib.Database{DB: tx}
	return NewAccountRepository(database)
}

func (r *accountRepositoryImpl) CheckBalance(ctx context.Context, accountNumber int64) (data models.CheckBalanceAccount, err error) {
	err = r.DB.Table("accounts a").
		Joins("LEFT JOIN customers c ON c.customer_number = a.customer_number").
		Select(
			"account_number",
			"customers.name AS customer_name",
			"balance",
		).
		Where("account_number = ?", accountNumber).
		Limit(1).
		WithContext(ctx).
		Find(&data).
		Error
	return
}

func (r *accountRepositoryImpl) GetAccountByPks(accountNumbers []*int64) (data []models.Account, err error) {
	err = r.DB.Raw(`SELECT * FROM accounts WHERE account_number IN ?;`, accountNumbers).Scan(&data).Error
	return
}

func (r *accountRepositoryImpl) CheckInsufficientBalance(accountNumber, amount *int64) (data models.Account, err error) {
	err = r.DB.Raw(`SELECT * FROM accounts WHERE account_number = ? AND balance >= ?`, accountNumber, amount).Scan(&data).Error
	return
}

func (r *accountRepositoryImpl) UpdateBalance(params *models.UpdateBalance) (err error) {
	err = r.DB.Exec(`
		UPDATE accounts SET balance = ((SELECT balance FROM accounts WHERE account_number = ?) + ?) WHERE account_number = ?;`,
		params.AccountNumber, params.Amount, params.AccountNumber,
	).Error

	return
}
