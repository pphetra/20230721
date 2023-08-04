package member_event_handlers

import (
	member_app_commands "taejai/internal/member/app/commands"
	member_domain_event "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"
)

type GreetingMailSendHandler struct {
	unitOfWork shared_app.UnitOfWork
}

func NewGreetingMailSendHandler(uow shared_app.UnitOfWork) *GreetingMailSendHandler {
	return &GreetingMailSendHandler{
		unitOfWork: uow,
	}
}

// implement EventHandler interface
func (h *GreetingMailSendHandler) GetEventName() string {
	return member_domain_event.GreetingMailSendEventName
}

func (h *GreetingMailSendHandler) Handle(event shared_domain.DomainEvent) error {
	sendMailEvent := event.(member_domain_event.GreetingMailSendEvent)

	// new updateMailSendCommand
	cmd := member_app_commands.UpdateMailSendCommand{
		MemberId: sendMailEvent.MemberId,
		Time:     sendMailEvent.GetOccuredAt(),
	}

	_, err := h.unitOfWork.GetCommandExecutor().Execute(cmd)
	return err
}
