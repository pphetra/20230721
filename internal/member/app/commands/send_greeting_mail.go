package member_app_commands

import (
	"errors"
	member_domain "taejai/internal/member/domain"
	member_domain_event "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
)

const SendGreetingMailCommandName = "SendGreetingMailCommand"

type SendGreetingMailCommand struct {
	MemberId member_domain.MemberId
	FullName string
	Email    string
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

	mailService, found := store.GetService("mail_service")
	if !found {
		return nil, errors.New("mail_service not found")
	}
	err := mailService.(shared_app.MailService).SendMail(to, subject, body)
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
