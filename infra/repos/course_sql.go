package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type courseSQLRepo struct {
	db *gorm.DB
}

func NewCourseSQLRepo(db *gorm.DB) *courseSQLRepo {
	return &courseSQLRepo{db: db}
}

func (r *courseSQLRepo) Create(ctx context.Context, c *models.Course) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *courseSQLRepo) GetByID(ctx context.Context, id string) (*models.Course, error) {
	var course models.Course
	if err := r.db.WithContext(ctx).First(&course, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseSQLRepo) List(ctx context.Context, fields map[string]interface{}, pagination *models.Pagination) ([]*models.Course, *models.Pagination, error) {
	var courses []*models.Course
	query := r.db.WithContext(ctx).Model(&models.Course{})

	for k, v := range fields {
		switch k {
		case "name":
			if name, ok := v.(string); ok {
				query = query.Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", "%"+name+"%")
			}
		case "session_count":
			query = query.Where("session_count = ?", v)
		case "price_min":
			query = query.Where("price_min >= ?", v)
		case "price_max":
			query = query.Where("price_max <= ?", v)
		case "order_by":
			switch v {
			case "price_asc":
				query = query.Order("price_per_session ASC")
			case "price_desc":
				query = query.Order("price_per_session DESC")
			case "session_count_asc":
				query = query.Order("session_count ASC")
			case "session_count_desc":
				query = query.Order("session_count DESC")
			}
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	pagination.SetTotal(total)
	query = pagination.ApplyToQuery(query)

	if err := query.Find(&courses).Error; err != nil {
		return nil, nil, err
	}

	return courses, pagination, nil
}

func (r *courseSQLRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&models.Course{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

func (r *courseSQLRepo) UpdateWithMap(ctx context.Context, id string, fields map[string]interface{}) (*models.Course, error) {
	course := &models.Course{}

	if err := r.db.WithContext(ctx).
		Model(course).
		Where("id = ?", id).
		Updates(fields).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		First(course, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return course, nil
}
