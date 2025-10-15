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

func (r *supplySQLRepo) GetByID(ctx context.Context, id string) (*models.Supply, error) {
	var supply models.Supply
	if err := r.db.WithContext(ctx).First(&supply, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &supply, nil
}

func (r *supplySQLRepo) List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.Supply, *models.Pagination, error) {
	var supplies []*models.Supply
	query := r.db.WithContext(ctx).Model(&models.Supply{})

	for k, v := range queries {
		switch k {
		case "name":
			query = query.Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", "%"+v.(string)+"%")
		case "min_threshold":
			query = query.Where("min_threshold >= ?", v)
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)
	query = pagination.ApplyToQuery(query)

	if err := query.Find(&supplies).Error; err != nil {
		return nil, nil, err
	}
	return supplies, pagination, nil
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
