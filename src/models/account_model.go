package models

type Account struct {
	AccountNumber  int64 `gorm:"type:integer" json:"account_number" validate:"required,number"`
	CustomerNumber int64 `gorm:"type:integer" json:"customer_number" validate:"required,number"`
	Balance        int64 `gorm:"integer" json:"balance" validate:"omitempty"`
}

type CheckSaldoAccount struct {
	AccountNumber int64  `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int64  `json:"balance"`
}
