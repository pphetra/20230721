package infra_eventbus_watermill

import (
	"context"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type WatermillGoChannelEventBus struct {
	publisher       message.Publisher
	subscriber      message.Subscriber
	commandExecutor *shared_app.CommandExecutor
}

func NewWatermillGoChannelEventBus(publisher message.Publisher, subscriber message.Subscriber) *WatermillGoChannelEventBus {
	return &WatermillGoChannelEventBus{
		publisher,
		subscriber,
		nil,
	}
}

func NewGoChannelEventBus() *WatermillGoChannelEventBus {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	return NewWatermillGoChannelEventBus(pubSub, pubSub)
}

func (b *WatermillGoChannelEventBus) SetCommandExecutor(commandExecutor *shared_app.CommandExecutor) {
	b.commandExecutor = commandExecutor
}

func (b *WatermillGoChannelEventBus) Publish(event shared_domain.DomainEvent) error {
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

	return b.publisher.Publish(event.GetEventName(), msg)
}

func (b *WatermillGoChannelEventBus) Subscribe(handler shared_app.EventHandler) error {
	messages, err := b.subscriber.Subscribe(context.Background(), handler.GetEventName())
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
				// TODO handle error
				continue
			}

			err = handler.Handle(b.commandExecutor, event)
			if err != nil {
				// TODO handle error
				continue
			}

			msg.Ack()
		}
	}()

	return nil
}
