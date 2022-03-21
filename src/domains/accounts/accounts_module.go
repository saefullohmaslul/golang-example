package accounts

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAccountService),
	fx.Provide(NewAccountRest),
	fx.Provide(NewAccountRoute),
)
