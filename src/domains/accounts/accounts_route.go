package accounts

import "restapi/src/lib"

type AccountRoute interface {
	Setup()
}

type AccountRouteImpl struct {
	echo lib.EchoHandler
	rest AccountRest
}

func NewAccountRoute(echo lib.EchoHandler, rest AccountRest) AccountRoute {
	return &AccountRouteImpl{
		echo: echo,
		rest: rest,
	}
}

func (r *AccountRouteImpl) Setup() {
	account := r.echo.Echo.Group("/account")
	{
		account.GET("/:account_number", r.rest.CheckBalance)
		account.POST("/account/:from_account_number/transfer", r.rest.Transfer)
	}
}
