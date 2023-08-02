package infra_mail

import (
	"fmt"
	shared_app "taejai/internal/shared/app"
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

func init() {
	shared_app.ServiceRegistry.Register("mail_service", NewNoOPMailService())
}
