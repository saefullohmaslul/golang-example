package interfaces

type EmailService interface {
	SendNotificationTransfer() (err error)
	SetAccountService(accountService AccountService)
}
