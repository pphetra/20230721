package infra_event_bus

import (
	"context"
	shared_app "taejai/internal/shared/app"
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

func (b *ChannelEventBus) Subscribe(commandExecutor *shared_app.CommandExecutor, handler shared_app.EventHandler) error {
	messages, err := b.pubSub.Subscribe(context.Background(), handler.GetEventName())
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			// deserialize event from msg.payload
			// get deserialized function from domain_event_registry
			deserializeFn := shared_domain.DomainEventReistry.GetDomainEventDeSerializerFunction(handler.GetEventName())
			event, err := deserializeFn(msg.Payload)
			if err != nil {
				// TODO log error
				continue
			}

			handler.Handle(commandExecutor, event)

			msg.Ack()
		}
	}()

	return nil
}
