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
