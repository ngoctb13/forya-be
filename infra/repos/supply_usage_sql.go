package repos

import (
	"context"
	"fmt"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type supplyUsageSQLRepo struct {
	db *gorm.DB
}

func NewSupplyUsageSQLRepo(db *gorm.DB) *supplyUsageSQLRepo {
	return &supplyUsageSQLRepo{
		db: db,
	}
}

func (r *supplyUsageSQLRepo) CreateUsagesAndDecreaseStock(ctx context.Context, usages []*models.SupplyUsage, adjustments map[string]int) error {
	if len(usages) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for batchID, qty := range adjustments {
			if qty == 0 {
				continue
			}

			res := tx.Model(&models.SupplyBatch{}).
				Where("id = ? AND remaining_quantity >= ?", batchID, qty).
				Update("remaining_quantity", gorm.Expr("remaining_quantity - ?", qty))

			if res.Error != nil {
				return res.Error
			}
			if res.RowsAffected == 0 {
				return fmt.Errorf("insufficient stock for batch %s", batchID)
			}
		}

		if err := tx.Create(usages).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *supplyUsageSQLRepo) RollbackUsages(ctx context.Context, usages []*models.SupplyUsage, adjustments map[string]int) error {
	if len(usages) == 0 {
		return nil
	}

	usageIDs := make([]string, 0, len(usages))
	for _, usage := range usages {
		if usage.ID != "" {
			usageIDs = append(usageIDs, usage.ID)
		}
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for batchID, qty := range adjustments {
			if qty == 0 {
				continue
			}
			if err := tx.Model(&models.SupplyBatch{}).
				Where("id = ?", batchID).
				Update("remaining_quantity", gorm.Expr("remaining_quantity + ?", qty)).Error; err != nil {
				return err
			}
		}

		if len(usageIDs) > 0 {
			if err := tx.Where("id IN ?", usageIDs).Delete(&models.SupplyUsage{}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
