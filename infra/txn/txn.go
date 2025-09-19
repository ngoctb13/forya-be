package txn

import (
	"context"

	"gorm.io/gorm"
)

type Txn struct {
	db *gorm.DB
}

func NewTxn(db *gorm.DB) ITxn {
	return &Txn{
		db: db,
	}
}

func (t *Txn) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
