package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ISupplyUsage interface {
	CreateUsagesAndDecreaseStock(ctx context.Context, usages []*models.SupplyUsage, adjustments map[string]int) error
	RollbackUsages(ctx context.Context, usages []*models.SupplyUsage, adjustments map[string]int) error
}
