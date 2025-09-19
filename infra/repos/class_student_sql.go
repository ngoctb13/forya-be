package repos

import (
	"context"
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type classStudentSQLRepo struct {
	db *gorm.DB
}

func NewClassStudentSQLRepo(db *gorm.DB) *classStudentSQLRepo {
	return &classStudentSQLRepo{
		db: db,
	}
}

func (r *classStudentSQLRepo) Create(ctx context.Context, cs *models.ClassStudent) error {
	return r.db.WithContext(ctx).Create(cs).Error
}

func (r *classStudentSQLRepo) BatchCreate(ctx context.Context, cs []*models.ClassStudent) error {
	return r.db.WithContext(ctx).Create(cs).Error
}

func (r *classStudentSQLRepo) MarkLeft(ctx context.Context, classID, studentID string) error {
	return r.db.WithContext(ctx).
		Model(&models.ClassStudent{}).
		Where("class_id = ? AND student_id = ?", classID, studentID).
		Update("left_at", time.Now()).Error
}
