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
			"c.name AS customer_name",
			"balance",
		).
		Where("account_number = ?", accountNumber).
		Limit(1).
		WithContext(ctx).
		Find(&data).
		Error
	return
}

func (r *accountRepositoryImpl) GetAccountByPks(ctx context.Context, accountNumbers []int64) (data []models.Account, err error) {
	err = r.DB.Table("accounts").
		Select("account_number", "customer_number", "balance").
		Where("account_number IN (?)", accountNumbers).
		Find(&data).
		WithContext(ctx).
		Error

	return
}

func (r *accountRepositoryImpl) CheckInsufficientBalance(ctx context.Context, accountNumber, amount int64) (data models.Account, err error) {
	err = r.DB.Table("accounts").
		Select("account_number", "customer_number", "balance").
		Where("account_number = ? AND balance >= ?", accountNumber, amount).
		Find(&data).
		WithContext(ctx).
		Error

	return
}

func (r *accountRepositoryImpl) UpdateBalance(ctx context.Context, params *models.UpdateBalance) (err error) {
	err = r.DB.Table("accounts").
		Where("account_number = ?", params.AccountNumber).
		Update("balance", r.DB.Table("accounts").Select("SUM(balance + ?)", params.Amount).Where("account_number = ?", params.AccountNumber).WithContext(ctx)).
		WithContext(ctx).
		Error

	return
}
