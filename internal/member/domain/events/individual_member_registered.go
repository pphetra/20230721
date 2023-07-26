package member_domain_event

import (
	"encoding/json"
	"fmt"
	member_domain "taejai/internal/member/domain"
	shared_domain "taejai/internal/shared/domain"
	"time"
)

const IndividualMemberRegisteredEventName = "individual_member_registered"

type IndividualMemberRegisteredEvent struct {
	MemberId          member_domain.MemberId `json:"member_id"`
	FirstName         string                 `json:"first_name"`
	LastName          string                 `json:"last_name"`
	Email             string                 `json:"email"`
	AddressLine1      string                 `json:"address_line_1"`
	AddressLine2      string                 `json:"address_line_2"`
	AddressPostalCode string                 `json:"address_postal_code"`
	shared_domain.BaseDomainEvent
}

func NewIndividualMemberRegisteredEvent(member member_domain.Member) IndividualMemberRegisteredEvent {
	return IndividualMemberRegisteredEvent{
		MemberId:          member.Id,
		FirstName:         member.Name1,
		LastName:          member.Name2,
		Email:             member.Email,
		AddressLine1:      member.Address.Line1,
		AddressLine2:      member.Address.Line2,
		AddressPostalCode: member.Address.PostalCode,
		BaseDomainEvent: shared_domain.BaseDomainEvent{
			OccuredAt: time.Now().Unix(),
		},
	}
}

// implement domain_event interface
func (e IndividualMemberRegisteredEvent) GetEventName() string {
	return IndividualMemberRegisteredEventName
}

func (e IndividualMemberRegisteredEvent) GetEventId() string {
	return fmt.Sprintf("member:%d", e.MemberId)
}

func (e IndividualMemberRegisteredEvent) GetOccuredAt() time.Time {
	return time.Unix(e.OccuredAt, 0)
}

func (e IndividualMemberRegisteredEvent) GetPayload() ([]byte, error) {
	return json.Marshal(e)
}

func (e IndividualMemberRegisteredEvent) GetAggregateId() string {
	return fmt.Sprintf("%d", e.MemberId)
}

func (e IndividualMemberRegisteredEvent) GetAggregateName() string {
	return "Member"
}

// ==============================================
func init() {
	shared_domain.DomainEventReistry.RegisterDomainEventDeSerializerFunction(IndividualMemberRegisteredEventName, fromPayload)
}

func fromPayload(payload []byte) (shared_domain.DomainEvent, error) {
	var event IndividualMemberRegisteredEvent
	err := json.Unmarshal(payload, &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
