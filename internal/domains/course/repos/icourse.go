package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ICourseRepo interface {
	Create(ctx context.Context, c *models.Course) error
	GetByID(ctx context.Context, id string) (*models.Course, error)
	List(ctx context.Context, fields map[string]interface{}, pagination *models.Pagination) ([]*models.Course, *models.Pagination, error)
	UpdateWithMap(ctx context.Context, id string, fields map[string]interface{}) (*models.Course, error)
	Delete(ctx context.Context, id string) error
}
