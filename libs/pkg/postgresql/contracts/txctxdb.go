package dbcontracts

import (
	"context"

	"gorm.io/gorm"
)

type TxContextDb interface {
	WithTx(ctx context.Context) (TxContextDb, error)
	WithTxIfExists(ctx context.Context) TxContextDb
	RunInTx(ctx context.Context, action ActionFunc) error
	DB() *gorm.DB
}

type ActionFunc func(ctx context.Context, txCtx TxContextDb) error
