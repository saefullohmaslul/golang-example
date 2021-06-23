package repositories

import "github.com/sarulabs/di"

func NewRepositoryMock(ioc di.Container) *Repository {
	return &Repository{}
}
