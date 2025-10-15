package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ISupply interface {
	Create(ctx context.Context, supply *models.Supply) error
	GetByID(ctx context.Context, supplyID string) (*models.Supply, error)
	List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.Supply, *models.Pagination, error)
	Delete(ctx context.Context, id string) error
	UpdateWithFields(ctx context.Context, supply *models.Supply, fields map[string]interface{}) error
}
