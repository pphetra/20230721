package member_app_commands

import (
	member_domain "taejai/internal/member/domain"
	member_domain_event "taejai/internal/member/domain/events"
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
	store.AddEventToPublish(
		member_domain_event.NewGreetingMailSendEvent(
			c.MemberId,
		),
	)

	return nil, nil
}
