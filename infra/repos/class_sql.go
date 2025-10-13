package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type classSQLRepo struct {
	db *gorm.DB
}

func NewClassSQLRepo(db *gorm.DB) *classSQLRepo {
	return &classSQLRepo{
		db: db,
	}
}

func (c *classSQLRepo) CreateClass(ctx context.Context, class *models.Class) error {
	return c.db.WithContext(ctx).Create(class).Error
}

func (c *classSQLRepo) GetClassByID(ctx context.Context, id string) (*models.Class, error) {
	var class models.Class
	err := c.db.WithContext(ctx).First(&class, "id = ?", id).Error
	return &class, err
}

func (c *classSQLRepo) GetClassContainName(ctx context.Context, name *string, pagination *models.Pagination) ([]*models.Class, *models.Pagination, error) {
	var classes []*models.Class
	query := c.db.WithContext(ctx).Model(&models.Class{})

	if name != nil {
		like := "%" + *name + "%"
		query = query.Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", like)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)

	query = pagination.ApplyToQuery(query)
	if err := query.Find(&classes).Error; err != nil {
		return nil, nil, err
	}

	return classes, pagination, nil
}
