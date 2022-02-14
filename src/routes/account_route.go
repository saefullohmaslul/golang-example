package routes

import (
	"restapi/src/controllers"
	"restapi/src/lib"
)

type AccountRouter struct {
	echoHandler lib.EchoHandler
	handler     controllers.AccountController
}

func (r *AccountRouter) Setup() {
	account := r.echoHandler.Echo.Group("/account")
	{
		account.GET("/:account_number", r.handler.CheckBalance)
		account.POST("/account/:from_account_number/transfer", r.handler.Transfer)
	}
}

func NewAccountRouter(
	handler controllers.AccountController,
	echoHandler lib.EchoHandler,
) AccountRouter {
	return AccountRouter{
		echoHandler: echoHandler,
		handler:     handler,
	}
}
