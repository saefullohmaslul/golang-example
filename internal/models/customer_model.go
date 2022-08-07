package models

type Customer struct {
	CustomerNumber int64  `gorm:"type:integer" json:"customer_number" validate:"required,number"`
	Name           string `gorm:"type:varchar(255)" json:"name" validate:"required"`
}
