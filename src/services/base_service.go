package services

import "github.com/sarulabs/di"

type Service struct {
	Account AccountService
}

func NewService(ioc di.Container) *Service {
	return &Service{
		Account: NewAccountService(ioc),
	}
}
