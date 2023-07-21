package shared_app

import shared_domain "taejai/internal/shared/domain"

type EventBus interface {
	Publish(event shared_domain.DomainEvent) error
}
