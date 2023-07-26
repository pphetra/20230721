package shared_app

import shared_domain "taejai/internal/shared/domain"

type EventBus interface {
	Publish(event shared_domain.DomainEvent) error
	Subscribe(handler EventHandler) error
}

type EventHandler interface {
	GetEventName() string
	Handle(commandExecutor *CommandExecutor, event shared_domain.DomainEvent) error
}
