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

func (r *classStudentSQLRepo) GetByClassAndStudent(ctx context.Context, classID, studentID string) (*models.ClassStudent, error) {
	var cs models.ClassStudent
	err := r.db.WithContext(ctx).
		Where("class_id = ? AND student_id = ?", classID, studentID).
		First(&cs).Error
	if err != nil {
		return nil, err
	}

	return &cs, nil
}

func (r *classStudentSQLRepo) ListByClassAndStudents(ctx context.Context, classID string, studentIDs []string) ([]*models.ClassStudent, error) {
	if len(studentIDs) == 0 {
		return []*models.ClassStudent{}, nil
	}

	var result []*models.ClassStudent
	if err := r.db.WithContext(ctx).
		Where("class_id = ? AND student_id IN ?", classID, studentIDs).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *classStudentSQLRepo) ResetLeftAtBulk(ctx context.Context, classID string, studentIDs []string) error {
	if len(studentIDs) == 0 {
		return nil
	}

	return r.db.WithContext(ctx).
		Model(&models.ClassStudent{}).
		Where("class_id = ? AND student_id IN ?", classID, studentIDs).
		Update("left_at", time.Time{}).Error
}
