package dbhelper

import (
	"context"
	"errors"

	dbconstants "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/constants"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"gorm.io/gorm"
)

func GetTxFromContext(ctx context.Context) (*gorm.DB, error) {
	gCtx, gCtxOk := ctx.(*dbcontracts.TxContext)
	if gCtxOk {
		return gCtx.Tx, nil
	}

	tx, ok := ctx.Value(dbconstants.TxKey).(*gorm.DB)
	if !ok {
		return nil, errors.New("Transaction not found in context")
	}

	return tx, nil
}

func GetTxFromContextIfExists(ctx context.Context) *gorm.DB {
	gCtx, gCtxOk := ctx.(*dbcontracts.TxContext)
	if gCtxOk {
		return gCtx.Tx
	}

	tx, ok := ctx.Value(dbconstants.TxKey).(*gorm.DB)
	if !ok {
		return nil
	}

	return tx
}

func SetTxToContext(ctx context.Context, tx *gorm.DB) *dbcontracts.TxContext {
	newCtx := context.WithValue(ctx, dbconstants.TxKey, tx)
	gormContext := &dbcontracts.TxContext{Tx: tx, Context: newCtx}
	ctx = gormContext

	return gormContext
}
