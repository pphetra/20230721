package infra_unit_of_work

import (
	"database/sql"
	shared_app "taejai/internal/shared/app"
)

type UnitOfWork struct {
	db              *sql.DB
	eventBus        shared_app.EventBus
	commandExecutor *shared_app.CommandExecutor
}

func NewUnitOfWork(db *sql.DB, eventBus shared_app.EventBus) *UnitOfWork {
	uow := &UnitOfWork{db: db, eventBus: eventBus}
	uow.commandExecutor = shared_app.NewCommandExecutor(uow)
	return uow
}

func (u *UnitOfWork) DoInTransaction(txFunc shared_app.UnitOfWorkTxFunc) (interface{}, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	uowStore := NewUnitOfWorkStore(tx)

	result, err := txFunc(uowStore)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	for _, event := range uowStore.EventToPublish {
		// publish event
		_ = u.eventBus.Publish(event)
		// TODO log error
	}

	return result, nil
}

func (u *UnitOfWork) GetCommandExecutor() *shared_app.CommandExecutor {
	return u.commandExecutor
}
