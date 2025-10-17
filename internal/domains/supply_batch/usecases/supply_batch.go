package usecases

import (
	"context"
	"fmt"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	supplyRp "github.com/ngoctb13/forya-be/internal/domains/supply/repos"
	"github.com/ngoctb13/forya-be/internal/domains/supply_batch/repos"
	"github.com/shopspring/decimal"
)

type SupplyBatch struct {
	supplyBatchRepo repos.ISupplyBatch
	supplyRepo      supplyRp.ISupply
}

func NewSupply(supplyBatchRepo repos.ISupplyBatch, supplyRepo supplyRp.ISupply) *SupplyBatch {
	return &SupplyBatch{
		supplyBatchRepo: supplyBatchRepo,
		supplyRepo:      supplyRepo,
	}
}

func (b *SupplyBatch) CreateSupplyBatch(ctx context.Context, input *inputs.CreateSupplyBatchInput) error {
	s, err := b.supplyRepo.GetByID(ctx, input.SupplyID)
	if err != nil {
		return err
	}
	if s == nil {
		return fmt.Errorf("supply not found")
	}

	sb := &models.SupplyBatch{
		SupplyID:          input.SupplyID,
		Quantity:          input.Quantity,
		RemainingQuantity: input.RemainingQuantity,
		PurchasePrice:     decimal.NewFromFloat(input.PurchasePrice),
		PurchaseDate:      input.PurchaseDate,
		Contact:           input.Contact,
	}

	return b.supplyBatchRepo.Create(ctx, sb)
}
