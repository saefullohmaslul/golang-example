package repositories

import (
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type Repository struct {
	Account AccountRepository
}

func NewRepository(ioc di.Container) *Repository {
	return &Repository{
		Account: NewAccountReposiory(ioc),
	}
}

func getDatabase(ioc di.Container) *gorm.DB {
	return ioc.Get("database").(*gorm.DB)
}
