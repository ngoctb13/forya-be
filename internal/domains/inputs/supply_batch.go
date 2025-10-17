package inputs

import (
	"time"
)

type CreateSupplyBatchInput struct {
	SupplyID          string
	Quantity          int
	RemainingQuantity int
	PurchasePrice     float64
	PurchaseDate      time.Time
	Contact           string
}
