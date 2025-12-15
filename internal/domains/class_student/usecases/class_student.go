package usecases

import (
	"context"
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/class_student/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type ClassStudent struct {
	classStudentRepo repos.IClassStudentRepo
}

func NewClassStudent(classStudentRepo repos.IClassStudentRepo) *ClassStudent {
	return &ClassStudent{
		classStudentRepo: classStudentRepo,
	}
}

func (c *ClassStudent) EnrollClass(ctx context.Context, input *inputs.EnrollClassInput) error {
	var csArr []*models.ClassStudent
	now := time.Now()

	existingArr, err := c.classStudentRepo.ListByClassAndStudents(ctx, input.ClassID, input.StudentIDs)
	if err != nil {
		return err
	}

	existingMap := make(map[string]*models.ClassStudent, len(existingArr))
	for _, cs := range existingArr {
		existingMap[cs.StudentID] = cs
	}

	var resetIDs []string

	for _, id := range input.StudentIDs {
		if cs, ok := existingMap[id]; ok {
			if cs.LeftAt.IsZero() {
				continue
			}
			resetIDs = append(resetIDs, id)
			continue
		}

		csArr = append(csArr, &models.ClassStudent{
			ClassID:   input.ClassID,
			StudentID: id,
			JoinedAt:  now,
		})
	}

	if len(resetIDs) > 0 {
		if err := c.classStudentRepo.ResetLeftAtBulk(ctx, input.ClassID, resetIDs); err != nil {
			return err
		}
	}

	if len(csArr) == 0 {
		return nil
	}

	return c.classStudentRepo.BatchCreate(ctx, csArr)
}

func (c *ClassStudent) DeleteStudentFromClass(ctx context.Context, classID, studentID string) error {
	return c.classStudentRepo.MarkLeft(ctx, classID, studentID)
}
