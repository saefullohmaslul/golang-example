package interfaces

import (
	"restapi/internal/lib"

	"gorm.io/gorm"
)

type Repository interface {
	Transaction(f func(tx *gorm.DB) error) error
}

type RepositoryImpl struct {
	lib.Database
}

func (r *RepositoryImpl) Transaction(f func(tx *gorm.DB) error) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		return f(tx)
	})
}
