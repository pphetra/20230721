package infra_event_bus

import (
	shared_domain "taejai/internal/shared/domain"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type ChannelEventBus struct {
	pubSub *gochannel.GoChannel
}

func NewChannelEventBus() *ChannelEventBus {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	return &ChannelEventBus{
		pubSub,
	}
}

func (b *ChannelEventBus) Publish(event shared_domain.DomainEvent) error {
	payload, err := event.GetPayload()
	if err != nil {
		return err
	}
	msg := message.NewMessage(
		event.GetAggregateId(),
		payload,
	)
	msg.Metadata.Set("aggregate_name", event.GetAggregateName())
	msg.Metadata.Set("aggregate_id", event.GetAggregateId())

	return b.pubSub.Publish(event.GetEventName(), msg)
}
