package repositories

import (
	"restapi/src/models"
	"restapi/src/utils/databases"

	"gorm.io/gorm"
)

type AccountRepository interface {
	CheckBalance(*int64) (models.CheckBalanceAccount, error)
	GetAccountByPks([]*int64) ([]models.Account, error)
	CheckInsufficientBalance(*int64, *int64) (models.Account, error)
	TransferBalance(*models.TransferBalance) error
}

type AccountRepositoryImpl struct {
	databases.Database
}

func NewAccountReposiory(db databases.Database) AccountRepository {
	return &AccountRepositoryImpl{
		Database: db,
	}
}

func (r *AccountRepositoryImpl) CheckBalance(accountNumber *int64) (data models.CheckBalanceAccount, err error) {
	err = r.DB.Raw(`
		SELECT account_number, customers.name AS customer_name, balance 
		FROM accounts LEFT JOIN customers ON customers.customer_number = accounts.customer_number
		WHERE account_number = ?;
	`, accountNumber).Scan(&data).Error
	return
}

func (r *AccountRepositoryImpl) GetAccountByPks(accountNumbers []*int64) (data []models.Account, err error) {
	err = r.DB.Raw(`SELECT * FROM accounts WHERE account_number IN ?;`, accountNumbers).Scan(&data).Error
	return
}

func (r *AccountRepositoryImpl) CheckInsufficientBalance(accountNumber, amount *int64) (data models.Account, err error) {
	err = r.DB.Raw(`SELECT * FROM accounts WHERE account_number = ? AND balance >= ?`, accountNumber, amount).Scan(&data).Error
	return
}

func (r *AccountRepositoryImpl) TransferBalance(bodies *models.TransferBalance) error {
	return r.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Exec(`
			UPDATE accounts SET balance = ((SELECT balance FROM accounts WHERE account_number = ?) + ?) WHERE account_number = ?;`,
			bodies.ToAccountNumber, bodies.Amount, bodies.ToAccountNumber,
		).Error; err != nil {
			return err
		}

		if err = tx.Exec(`
			UPDATE accounts SET balance = ((SELECT balance FROM accounts WHERE account_number = ?) - ?) WHERE account_number = ?;`,
			bodies.FromAccountNumber, bodies.Amount, bodies.FromAccountNumber,
		).Error; err != nil {
			return err
		}

		return nil
	})
}
