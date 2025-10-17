package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ISupplyBatch interface {
	Create(ctx context.Context, supply *models.SupplyBatch) error
	GetByID(ctx context.Context, id string) (*models.SupplyBatch, error)
	List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.SupplyBatch, *models.Pagination, error)
	UpdateWithFields(ctx context.Context, sb *models.SupplyBatch, fields map[string]interface{}) error
}
