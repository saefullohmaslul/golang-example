package email

import (
	"fmt"
	"restapi/internal/interfaces"
)

type emailServiceImpl struct {
	accountService interfaces.AccountService
}

func NewEmailService() interfaces.EmailService {
	return &emailServiceImpl{}
}

func (s *emailServiceImpl) SetAccountService(accountService interfaces.AccountService) {
	// you can use other module
	s.accountService = accountService
}

func (s *emailServiceImpl) SendNotificationTransfer() (err error) {
	// TODO: this is implementation for send notification
	fmt.Println("success send email notification")
	return
}
