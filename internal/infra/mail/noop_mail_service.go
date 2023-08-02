package infra_mail

import (
	"fmt"
)

type NoOpMailService struct {
}

func NewNoOPMailService() *NoOpMailService {
	return &NoOpMailService{}
}

// implement MailService interface
func (s *NoOpMailService) SendMail(to string, subject string, body string) error {
	fmt.Println("SendMail  to" + to + " subject:" + subject + " body:" + body)
	return nil
}
