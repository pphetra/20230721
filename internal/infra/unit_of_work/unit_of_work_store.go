package infra_unit_of_work

import (
	"database/sql"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"
)

type UnitOfWorkStore struct {
	tx             *sql.Tx
	EventToPublish []shared_domain.DomainEvent
}

func NewUnitOfWorkStore(tx *sql.Tx) *UnitOfWorkStore {
	return &UnitOfWorkStore{tx: tx}
}

func (s *UnitOfWorkStore) GetRepository(name string) interface{} {
	return shared_domain.RepositoryRegistry.GetRepository(name, s.tx)
}

func (s *UnitOfWorkStore) AddEventToPublish(event shared_domain.DomainEvent) error {
	s.EventToPublish = append(s.EventToPublish, event)
	return nil
}

func (s *UnitOfWorkStore) GetService(name string) (interface{}, bool) {
	return shared_app.ServiceRegistry.GetService(name)
}
