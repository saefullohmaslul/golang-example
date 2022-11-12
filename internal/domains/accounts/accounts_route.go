package accounts

import (
	"restapi/internal/interfaces"
	"restapi/internal/lib"
)

type AccountRouteImpl struct {
	echo lib.EchoHandler
	rest interfaces.AccountRest
}

func NewAccountRoute(echo lib.EchoHandler, rest interfaces.AccountRest) interfaces.AccountRoute {
	return &AccountRouteImpl{
		echo: echo,
		rest: rest,
	}
}

func (r *AccountRouteImpl) Setup() {
	account := r.echo.Echo.Group("/account")
	{
		account.GET("/:account_number", r.rest.CheckBalance)
		account.POST("/:from_account_number/transfer", r.rest.Transfer)
	}
}
