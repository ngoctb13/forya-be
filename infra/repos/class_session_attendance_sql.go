package repos

import (
	"context"

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
