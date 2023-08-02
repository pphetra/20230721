package shared_app

import shared_domain "taejai/internal/shared/domain"

type UnitOfWorkStore interface {
	// return nil if not found
	GetRepository(name string) interface{}
	AddEventToPublish(event shared_domain.DomainEvent) error
	GetService(name string) (interface{}, bool)
}

type UnitOfWorkTxFunc func(store UnitOfWorkStore) (interface{}, error)

type UnitOfWork interface {
	DoInTransaction(UnitOfWorkTxFunc) (interface{}, error)
	GetCommandExecutor() *CommandExecutor
}
