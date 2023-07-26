package infra_mail

import "fmt"

type MailService struct {
}

func NewMailService() *MailService {
	return &MailService{}
}

// implement MailService interface
func (s *MailService) SendMail(to string, subject string, body string) error {
	fmt.Println("SendMail  to" + to + " subject:" + subject + " body:" + body)
	return nil
}
