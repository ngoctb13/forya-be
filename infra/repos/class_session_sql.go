package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type classSessionSQLRepo struct {
	db *gorm.DB
}

func NewClassSessionSQLRepo(db *gorm.DB) *classSessionSQLRepo {
	return &classSessionSQLRepo{
		db: db,
	}
}

func (r *classSessionSQLRepo) Create(ctx context.Context, session *models.ClassSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *classSessionSQLRepo) List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.ClassSession, *models.Pagination, error) {
	query := r.db.WithContext(ctx).Model(&models.ClassSession{})

	for k, v := range queries {
		switch k {
		case "class_id":
			query = query.Where("class_id = ?", v)
		case "start_time":
			query = query.Where("held_at >= ?", v)
		case "end_time":
			query = query.Where("held_at <= ?", v)
		}
	}

	var (
		total int64
		csArr []*models.ClassSession
	)

	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)
	query = pagination.ApplyToQuery(query)

	if err := query.Preload("classes").Order("held_at DESC").Find(&csArr).Error; err != nil {
		return nil, nil, err
	}

	return csArr, pagination, nil
}
