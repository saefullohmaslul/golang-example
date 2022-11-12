package email

import (
	"fmt"
	"restapi/internal/interfaces"
)

type emailServiceImpl struct{}

func NewEmailService() interfaces.EmailService {
	return &emailServiceImpl{}
}

func (s *emailServiceImpl) SendNotificationTransfer() (err error) {
	// TODO: this is implementation for send notification
	fmt.Println("success send email notification")
	return
}
