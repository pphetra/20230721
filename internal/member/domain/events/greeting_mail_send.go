package member_domain_event

import (
	"encoding/json"
	"fmt"
	member "taejai/internal/member/domain"
	shared_domain "taejai/internal/shared/domain"
	"time"
)

const GreetingMailSendEventName = "greeting_mail_send"

type GreetingMailSendEvent struct {
	MemberId member.MemberId `json:"member_id"`
	shared_domain.BaseDomainEvent
}

func NewGreetingMailSendEvent(memberId member.MemberId) GreetingMailSendEvent {
	return GreetingMailSendEvent{
		MemberId: memberId,
		BaseDomainEvent: shared_domain.BaseDomainEvent{
			OccuredAt: time.Now().Unix(),
		},
	}
}

// implement domain_event interface
func (e GreetingMailSendEvent) GetEventName() string {
	return GreetingMailSendEventName
}

func (e GreetingMailSendEvent) GetOccuredAt() time.Time {
	return time.Unix(e.OccuredAt, 0)
}

func (e GreetingMailSendEvent) GetPayload() ([]byte, error) {
	return json.Marshal(e)
}

func (e GreetingMailSendEvent) GetAggregateId() string {
	return fmt.Sprintf("%d", e.MemberId)
}

func (e GreetingMailSendEvent) GetAggregateName() string {
	return "member"
}

// ====================
func init() {
	shared_domain.DomainEventReistry.RegisterDomainEventDeSerializerFunction(GreetingMailSendEventName, parseGreetingMailSendEvent)
}

func parseGreetingMailSendEvent(payload []byte) (shared_domain.DomainEvent, error) {
	var event GreetingMailSendEvent
	err := json.Unmarshal(payload, &event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
