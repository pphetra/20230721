package infra_unit_of_work

import (
	"database/sql"
	shared_app "taejai/internal/shared/app"
)

type UnitOfWork struct {
	db       *sql.DB
	eventBus shared_app.EventBus
}

func NewUnitOfWork(db *sql.DB, eventBus shared_app.EventBus) *UnitOfWork {
	return &UnitOfWork{db: db, eventBus: eventBus}
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
