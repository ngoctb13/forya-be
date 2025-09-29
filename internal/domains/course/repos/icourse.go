package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ICourseRepo interface {
	Create(ctx context.Context, c *models.Course) error
	GetByID(ctx context.Context, id string) (*models.Course, error)
	GetAll(ctx context.Context, filter *models.GetAllFilter) ([]*models.Course, error)
	Update(ctx context.Context, c *models.Course) error
	UpdateWithMap(ctx context.Context, id string, fields map[string]interface{}) (*models.Course, error)
	Delete(ctx context.Context, id string) error
	SearchByName(ctx context.Context, keyword string) ([]*models.Course, error)
}
