package repos

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type classSessionAttendanceSQLRepo struct {
	db *gorm.DB
}

func NewClassSessionAttendanceSQLRepo(db *gorm.DB) *classSessionAttendanceSQLRepo {
	return &classSessionAttendanceSQLRepo{
		db: db,
	}
}

func (r *classSessionAttendanceSQLRepo) MarkAttendance(ctx context.Context, sessionID, courseStudentID string, isAttended bool) (*models.ClassSessionAttendance, error) {
	attendance := &models.ClassSessionAttendance{}

	err := r.db.WithContext(ctx).
		Where("class_session_id = ? AND course_student_id = ?", sessionID, courseStudentID).
		First(attendance).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		attendance = &models.ClassSessionAttendance{
			ClassSessionID:  sessionID,
			CourseStudentID: courseStudentID,
			IsAttended:      isAttended,
		}
		if err := r.db.WithContext(ctx).Create(attendance).Error; err != nil {
			return nil, err
		}
		return attendance, nil
	}
	if err != nil {
		return nil, err
	}

	attendance.IsAttended = isAttended
	if err := r.db.WithContext(ctx).Save(attendance).Error; err != nil {
		return nil, err
	}

	return attendance, nil
}
