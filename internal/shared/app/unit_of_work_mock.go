package shared_app

import (
	shared_domain "taejai/internal/shared/domain"

	"github.com/stretchr/testify/mock"
)

type MockUnitOfWorkStore struct {
	mock.Mock
}

func (m *MockUnitOfWorkStore) GetRepository(name string) interface{} {
	ret := m.Called(name)
	return ret.Get(0)
}

func (m *MockUnitOfWorkStore) AddEventToPublish(event shared_domain.DomainEvent) error {
	ret := m.Called(event)
	return ret.Error(0)
}

func (m *MockUnitOfWorkStore) GetService(name string) (interface{}, bool) {
	ret := m.Called(name)
	return ret.Get(0), ret.Bool(1)
}

type MockUnitOfWork struct {
	mock.Mock
	UnitOfWorkStore UnitOfWorkStore
}

func (m *MockUnitOfWork) DoInTransaction(fn UnitOfWorkTxFunc) (interface{}, error) {
	return fn(m.UnitOfWorkStore)
}
