package infra_unit_of_work

import (
	"database/sql"
	shared_app "taejai/internal/shared/app"
)

type PostgresUnitOfWork struct {
	db       *sql.DB
	eventBus shared_app.EventBus
}

func NewPostgresUnitOfWork(db *sql.DB, eventBus shared_app.EventBus) *PostgresUnitOfWork {
	return &PostgresUnitOfWork{db: db, eventBus: eventBus}
}

func (u *PostgresUnitOfWork) DoInTransaction(txFunc shared_app.UnitOfWorkTxFunc) (interface{}, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	uowStore := NewPostgresUnitOfWorkStore(tx)

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
