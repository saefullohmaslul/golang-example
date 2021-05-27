package controllers

import "github.com/sarulabs/di"

type Controller struct {
	Account AccountController
}

func NewController(ioc di.Container) *Controller {
	return &Controller{
		Account: NewAccountController(ioc),
	}
}
