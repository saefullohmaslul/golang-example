package fixtures

import "restapi/src/models"

var (
	CheckBalanceAccount = models.CheckBalanceAccount{
		AccountNumber: 1,
		CustomerName:  "Test",
		Balance:       1000,
	}
)

var (
	TransferBalance = models.TransferBalance{
		ToAccountNumber:   1,
		FromAccountNumber: 2,
		Amount:            1000,
	}

	Account = models.Account{
		AccountNumber:  2,
		CustomerNumber: 2,
		Balance:        10000,
	}
)
