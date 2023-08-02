package member_event_handlers

import (
	member_app_commands "taejai/internal/member/app/commands"
	member_domain_event "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"
)

type SendGreetingMailHandler struct {
}

func NewSendGreetingMailHandler() *SendGreetingMailHandler {
	return &SendGreetingMailHandler{}
}

// implement EventHandler interface
func (h *SendGreetingMailHandler) GetEventName() string {
	return member_domain_event.IndividualMemberRegisteredEventName
}

// implement EventHandler interface
func (h *SendGreetingMailHandler) Handle(commandExecutor *shared_app.CommandExecutor, event shared_domain.DomainEvent) error {
	registeredEvent := event.(member_domain_event.IndividualMemberRegisteredEvent)

	// new sendMailCommand
	cmd := member_app_commands.SendGreetingMailCommand{
		MemberId: registeredEvent.MemberId,
		FullName: registeredEvent.FirstName + " " + registeredEvent.LastName,
		Email:    registeredEvent.Email,
	}

	// execute sendMailCommand
	_, err := commandExecutor.Execute(cmd)
	if err != nil {
		return err
	}

	return nil
}
