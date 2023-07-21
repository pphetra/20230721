package shared_app

import (
	shared_domain "taejai/internal/shared/domain"

	"github.com/stretchr/testify/mock"
)

type MockEventBus struct {
	mock.Mock
}

func (m *MockEventBus) Publish(event shared_domain.DomainEvent) error {
	ret := m.Called(event)
	return ret.Error(0)
}
