package request

import (
	"errors"
	"time"

	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type CreateSupplyBatchRequest struct {
	SupplyID          string    `json:"supply_id"`
	Quantity          int       `json:"quantity"`
	RemainingQuantity int       `json:"remaining_quantity"`
	PurchasePrice     float64   `json:"purchase_price"`
	PurchaseDate      time.Time `json:"purchase_date"`
	Contact           string    `json:"contact"`
}

func (r *CreateSupplyBatchRequest) ValidateAndMap() (*inputs.CreateSupplyBatchInput, error) {
	if r.SupplyID == "" {
		return nil, errors.New("supply_id is required")
	}
	if r.Quantity <= 0 {
		return nil, errors.New("quantity must be > 0")
	}
	if r.RemainingQuantity < 0 || r.RemainingQuantity > r.Quantity {
		return nil, errors.New("remaining_quantity must be between 0 and quantity")
	}
	if r.PurchasePrice < 0 {
		return nil, errors.New("purchase_price must be >= 0")
	}

	input := &inputs.CreateSupplyBatchInput{
		SupplyID:          r.SupplyID,
		Quantity:          r.Quantity,
		RemainingQuantity: r.RemainingQuantity,
		PurchasePrice:     r.PurchasePrice,
		PurchaseDate:      r.PurchaseDate,
		Contact:           r.Contact,
	}
	return input, nil
}
