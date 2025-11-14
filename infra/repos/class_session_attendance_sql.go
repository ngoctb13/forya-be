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

func (r *classSessionAttendanceSQLRepo) BatchMarkAttendance(ctx context.Context, sessionID string, attendances []*models.ClassSessionAttendance) error {
	if len(attendances) == 0 {
		return nil
	}

	courseStudentIDs := make([]string, 0, len(attendances))
	for _, att := range attendances {
		att.ClassSessionID = sessionID
		courseStudentIDs = append(courseStudentIDs, att.CourseStudentID)
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingAttendances []*models.ClassSessionAttendance
		if err := tx.Where("class_session_id = ? AND course_student_id IN ?", sessionID, courseStudentIDs).
			Find(&existingAttendances).Error; err != nil {
			return err
		}

		existingMap := make(map[string]*models.ClassSessionAttendance, len(existingAttendances))
		for _, existing := range existingAttendances {
			existingMap[existing.CourseStudentID] = existing
		}

		var toCreate []*models.ClassSessionAttendance
		var toUpdate []*models.ClassSessionAttendance

		for _, att := range attendances {
			if existing, exists := existingMap[att.CourseStudentID]; exists {
				existing.IsAttended = att.IsAttended
				toUpdate = append(toUpdate, existing)
			} else {
				toCreate = append(toCreate, att)
			}
		}

		if len(toCreate) > 0 {
			if err := tx.Create(toCreate).Error; err != nil {
				return err
			}
		}

		if len(toUpdate) > 0 {
			if err := tx.Save(toUpdate).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
