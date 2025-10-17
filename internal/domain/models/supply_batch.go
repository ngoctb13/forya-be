package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type SupplyBatch struct {
	ID                string
	SupplyID          string
	Quantity          int
	RemainingQuantity int
	PurchasePrice     decimal.Decimal
	PurchaseDate      time.Time
	Contact           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
