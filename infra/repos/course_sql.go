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

func (r *courseSQLRepo) GetAll(ctx context.Context, filter *models.GetAllFilter) ([]*models.Course, error) {
	var courses []*models.Course
	q := r.db.WithContext(ctx).Model(&models.Course{})

	if filter.Name != nil {
		q = q.Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", "%"+*filter.Name+"%")
	}

	if filter.Description != nil {
		q = q.Where("unaccent(lower(description)) ILIKE unaccent(lower(?))", "%"+*filter.Description+"%")
	}

	if filter.SessionCount != nil {
		q = q.Where("session_count = ?", *filter.SessionCount)
	}

	if filter.PriceMin != nil {
		q = q.Where("price_per_session >= ?", *filter.PriceMin)
	}

	if filter.PriceMax != nil {
		q = q.Where("price_per_session <= ?", *filter.PriceMax)
	}

	if filter.OrderBy != nil {
		switch *filter.OrderBy {
		case "price_asc":
			q = q.Order("price_per_session ASC")
		case "price_desc":
			q = q.Order("price_per_session DESC")
		case "session_count_asc":
			q = q.Order("session_count ASC")
		case "session_count_desc":
			q = q.Order("session_count DESC")
		}
	}

	err := q.Find(&courses).Error
	return courses, err
}

func (r *courseSQLRepo) Update(ctx context.Context, c *models.Course) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *courseSQLRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&models.Course{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

func (r *courseSQLRepo) SearchByName(ctx context.Context, keyword string) ([]*models.Course, error) {
	var courses []*models.Course
	err := r.db.WithContext(ctx).
		Where("unaccent(lower(name)) ILIKE unaccent(lower(?))", "%"+keyword+"%").
		Find(&courses).Error
	return courses, err
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
