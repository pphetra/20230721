package member_app_commands

import (
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"
)

const SendGreetingMailCommandName = "SendGreetingMailCommand"

type SendGreetingMailCommand struct {
	MemberId    member_domain.MemberId
	FullName    string
	Email       string
	MailService shared_app.MailService
}

// implement Command interface
func (c SendGreetingMailCommand) GetCommandName() string {
	return SendGreetingMailCommandName
}

// implement Command interface
func (c SendGreetingMailCommand) Execute(store shared_app.UnitOfWorkStore) (interface{}, error) {
	to := c.Email
	subject := "Welcome to Taejai"
	body := "Dear " + c.FullName + ",\n\nWelcome to Taejai!\n\nBest Regards,\nTaejai Team"
	err := c.MailService.SendMail(to, subject, body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
