package postgres

import (
	"context"
	"log"
)

type TxManager struct {
	db *DB
}

func NewTxManager(db *DB) *TxManager {
	return &TxManager{db: db}
}

func (tm *TxManager) WithTx(ctx context.Context, fn func(txCtx context.Context) error) error {
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
