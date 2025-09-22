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

func (c *classSQLRepo) GetClassContainName(ctx context.Context, name *string) ([]*models.Class, error) {
	var classes []*models.Class
	query := c.db.WithContext(ctx).Debug().Model(&models.Class{})

	if name != nil {
		like := "%" + *name + "%"
		query = query.Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", like)
	}

	err := query.Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, err
}
