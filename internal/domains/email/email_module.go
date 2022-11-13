package email

import (
	"restapi/internal/interfaces"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewEmailService),
	fx.Invoke(func(accountService interfaces.AccountService, emailService interfaces.EmailService) {
		emailService.SetAccountService(accountService)
	}),
)
