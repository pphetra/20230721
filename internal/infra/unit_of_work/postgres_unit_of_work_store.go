package infra_unit_of_work

import (
	"database/sql"
	member_infra "taejai/internal/member/infra"
	shared_domain "taejai/internal/shared/domain"
)

type PostgresUnitOfWorkStore struct {
	tx             *sql.Tx
	EventToPublish []shared_domain.DomainEvent
}

func NewPostgresUnitOfWorkStore(tx *sql.Tx) *PostgresUnitOfWorkStore {
	return &PostgresUnitOfWorkStore{tx: tx}
}

func (s *PostgresUnitOfWorkStore) GetRepository(name string) interface{} {
	switch name {
	case "member":
		return member_infra.NewMemberRepository(s.tx)
	default:
		return nil
	}
}

func (s *PostgresUnitOfWorkStore) AddEventToPublish(event shared_domain.DomainEvent) error {
	s.EventToPublish = append(s.EventToPublish, event)
	return nil
}
