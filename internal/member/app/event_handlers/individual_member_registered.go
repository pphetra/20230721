package member_event_handlers

import (
	member_domain_event "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"
)

type IndividualMemberRegisteredHandler struct {
}

func NewIndividualMemberRegisteredHandler() *IndividualMemberRegisteredHandler {
	return &IndividualMemberRegisteredHandler{}
}

// implement EventHandler interface
func (h *IndividualMemberRegisteredHandler) GetEventName() string {
	return member_domain_event.IndividualMemberRegisteredEventName
}

// implement EventHandler interface
func (h *IndividualMemberRegisteredHandler) Handle(commandExecutor *shared_app.CommandExecutor, event shared_domain.DomainEvent) error {
	// new sendMailCommand

	// execute sendMailCommand

	return nil
}
