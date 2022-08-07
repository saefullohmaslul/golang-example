package repositories

import (
	"restapi/internal/models"
)

func (r *RepositoryImpl) CheckBalance(accountNumber *int64) (data models.CheckBalanceAccount, err error) {
	err = r.DB.Raw(`
		SELECT account_number, customers.name AS customer_name, balance 
		FROM accounts LEFT JOIN customers ON customers.customer_number = accounts.customer_number
		WHERE account_number = ?;
	`, accountNumber).Scan(&data).Error
	return
}

func (r *RepositoryImpl) GetAccountByPks(accountNumbers []*int64) (data []models.Account, err error) {
	err = r.DB.Raw(`SELECT * FROM accounts WHERE account_number IN ?;`, accountNumbers).Scan(&data).Error
	return
}

func (r *RepositoryImpl) CheckInsufficientBalance(accountNumber, amount *int64) (data models.Account, err error) {
	err = r.DB.Raw(`SELECT * FROM accounts WHERE account_number = ? AND balance >= ?`, accountNumber, amount).Scan(&data).Error
	return
}

func (r *RepositoryImpl) UpdateBalance(params *models.UpdateBalance) (err error) {
	err = r.DB.Exec(`
		UPDATE accounts SET balance = ((SELECT balance FROM accounts WHERE account_number = ?) + ?) WHERE account_number = ?;`,
		params.AccountNumber, params.Amount, params.AccountNumber,
	).Error

	return
}
