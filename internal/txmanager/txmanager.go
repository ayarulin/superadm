package txmanager

import (
	"context"
	"log"

	"superadmin.ru/internal/eventbus"
	"superadmin.ru/internal/postgres"
)

type TxManager struct {
	db       *postgres.DB
	eventBus *eventbus.SyncEventBus
}

type TxContext interface {
	context.Context
	AddEvent(e any)
}

func NewTxManager(db *postgres.DB, eventBus *eventbus.SyncEventBus) *TxManager {
	return &TxManager{
		db:       db,
		eventBus: eventBus,
	}
}

func (tm *TxManager) WithTx(ctx context.Context, fn func(txCtx TxContext) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)

	if err != nil {
		log.Panicf("begin tx: %v", err)
	}

	txCtx := newTxContext(ctx, tx)

	if err := fn(txCtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Panicf("rollback error: %v (original error: %v)", rbErr, err)
		}

		return err
	}

	return tx.Commit()
}
