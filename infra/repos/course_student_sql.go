package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type courseStudentSQLRepo struct {
	db *gorm.DB
}

func NewCourseStudentSQLRepo(db *gorm.DB) *courseStudentSQLRepo {
	return &courseStudentSQLRepo{
		db: db,
	}
}

func (r *courseStudentSQLRepo) Create(ctx context.Context, cs *models.CourseStudent) error {
	return r.db.WithContext(ctx).Create(cs).Error
}

func (r *courseStudentSQLRepo) BatchCreate(ctx context.Context, cs []*models.CourseStudent) error {
	return r.db.WithContext(ctx).Create(cs).Error
}

func (r *courseStudentSQLRepo) GetByIDs(ctx context.Context, ids []string) (map[string]*models.CourseStudent, error) {
	result := make(map[string]*models.CourseStudent)
	if len(ids) == 0 {
		return result, nil
	}

	var courseStudents []*models.CourseStudent
	if err := r.db.WithContext(ctx).
		Where("id IN ?", ids).
		Find(&courseStudents).Error; err != nil {
		return nil, err
	}

	for _, cs := range courseStudents {
		result[cs.ID] = cs
	}

	return result, nil
}
