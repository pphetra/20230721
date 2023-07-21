package shared_domain

import "time"

type DomainEvent interface {
	GetAggregateId() string
	GetAggregateName() string
	GetEventName() string
	GetOccuredAt() time.Time
	GetPayload() ([]byte, error)
}

type BaseDomainEvent struct {
	OccuredAt int64 `json:"occured_at"`
}

func (e BaseDomainEvent) GetOccuredAt() time.Time {
	return time.Unix(e.OccuredAt, 0)
}
