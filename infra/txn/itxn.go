package txn

import (
	"context"

	"gorm.io/gorm"
)

type ITxn interface {
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}
