package dbcontracts

import (
	"context"

	"gorm.io/gorm"
)

type TxContext struct {
	context.Context
	Tx *gorm.DB
}
