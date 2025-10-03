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
	return r.db.WithContext(ctx).Create(supply).Error
}

func (r *supplySQLRepo) ListByName(ctx context.Context, keyword string) ([]*models.Supply, error) {
	var supplies []*models.Supply
	err := r.db.WithContext(ctx).
		Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", "%"+keyword+"%").
		Find(&supplies).Error
	return supplies, err
}

func (r *supplySQLRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&models.Supply{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

func (r *supplySQLRepo) UpdateWithFields(ctx context.Context, supply *models.Supply, fields map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(supply).Updates(fields).Error
}
