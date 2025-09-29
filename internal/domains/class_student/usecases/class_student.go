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

	for _, id := range input.StudentIDs {
		cs := &models.ClassStudent{
			ClassID:   input.ClassID,
			StudentID: id,
			JoinedAt:  now,
		}

		csArr = append(csArr, cs)
	}

	return c.classStudentRepo.BatchCreate(ctx, csArr)
}

func (c *ClassStudent) DeleteStudentFromClass(ctx context.Context, classID, studentID string) error {
	return c.classStudentRepo.MarkLeft(ctx, classID, studentID)
}
