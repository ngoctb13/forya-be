package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassSession interface {
	Create(ctx context.Context, session *models.ClassSession) error
	GetByID(ctx context.Context, id string) (*models.ClassSession, error)
	List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.ClassSession, *models.Pagination, error)
}
