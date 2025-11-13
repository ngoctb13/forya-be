package repos

import (
	"context"
	"errors"

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

func (r *classSessionSQLRepo) GetByID(ctx context.Context, id string) (*models.ClassSession, error) {
	var session models.ClassSession
	if err := r.db.WithContext(ctx).
		First(&session, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func (r *classSessionSQLRepo) List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.ClassSession, *models.Pagination, error) {
	query := r.db.WithContext(ctx).Model(&models.ClassSession{})

	// Apply filters
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

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)

	// Apply pagination and ordering
	query = pagination.ApplyToQuery(query)
	query = query.Order("held_at DESC")

	// Fetch ClassSessions (Class will be populated in usecase via batch loading)
	if err := query.Find(&csArr).Error; err != nil {
		return nil, nil, err
	}

	return csArr, pagination, nil
}
