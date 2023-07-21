package member_domain_event_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	member_domain_event "taejai/internal/member/domain/events"
	shared_domain "taejai/internal/shared/domain"
)

func TestIndividualMemberRegisteredEvent(t *testing.T) {
	// Create a new event.
	event := member_domain_event.IndividualMemberRegisteredEvent{
		MemberId:          123,
		FirstName:         "John",
		LastName:          "Doe",
		AddressLine1:      "123 Main St",
		AddressLine2:      "Apt 4",
		AddressPostalCode: "12345",
	}
	event.OccuredAt = time.Now().Unix()

	// Serialize the event to JSON.
	payload, err := json.Marshal(event)
	assert.NoError(t, err)

	// Deserialize the event from JSON.
	deserializedEvent, err := shared_domain.DomainEventReistry.GetDomainEventDeSerializerFunction(member_domain_event.IndividualMemberRegisteredEventName)(payload)
	assert.NoError(t, err)

	// Ensure that the deserialized event is equal to the original event.
	assert.Equal(t, event, deserializedEvent)
}
