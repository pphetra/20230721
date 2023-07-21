package infra_event_bus

import (
	shared_domain "taejai/internal/shared/domain"
)

type ChannelEventBus struct {
}

func NewChannelEventBus() *ChannelEventBus {
	return &ChannelEventBus{}
}

func (b *ChannelEventBus) Publish(event shared_domain.DomainEvent) error {
	return nil
}
