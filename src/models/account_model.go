package models

type Account struct {
	AccountNumber  int64 `gorm:"type:integer" json:"account_number" validate:"required,number"`
	CustomerNumber int64 `gorm:"type:integer" json:"customer_number" validate:"required,number"`
	Balance        int64 `gorm:"integer" json:"balance" validate:"omitempty"`
}

type CheckBalanceAccount struct {
	AccountNumber int64  `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int64  `json:"balance"`
}

type TransferBalance struct {
	ToAccountNumber   int64 `json:"to_account_number" validate:"required,number"`
	Amount            int64 `json:"amount" validate:"required,number"`
	FromAccountNumber int64 `json:"from_account_number" validate:"required,number"`
}
