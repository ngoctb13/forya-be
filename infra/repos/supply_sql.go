package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type supplySQLRepo struct {
	db *gorm.DB
}

func NewSupplySQLRepo(db *gorm.DB) *supplySQLRepo {
	return &supplySQLRepo{db: db}
}

func (r *supplySQLRepo) Create(ctx context.Context, supply *models.Supply) error {
	return nil
}

func (r *supplySQLRepo) ListByName(ctx context.Context, keyword string) ([]*models.Supply, error) {
	return nil, nil
}

func (r *supplySQLRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *supplySQLRepo) UpdateWithFields(ctx context.Context, supply *models.Supply, fields map[string]interface{}) error {
	return nil
}
