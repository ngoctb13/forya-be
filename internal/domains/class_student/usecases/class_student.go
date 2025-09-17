package usecases

import (
	"context"
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/class_student/repos"
)

type ClassStudent struct {
	classStudentRepo repos.IClassStudentRepo
}

func NewClassStudent(classStudentRepo repos.IClassStudentRepo) *ClassStudent {
	return &ClassStudent{
		classStudentRepo: classStudentRepo,
	}
}

func (c *ClassStudent) EnrollStudent(ctx context.Context, input *models.EnrollStudentInput) error {
	cs := &models.ClassStudent{
		ClassID:   input.ClassID,
		StudentID: input.StudentID,
		JoinedAt:  time.Now(),
	}

	return c.classStudentRepo.Create(ctx, cs)
}

func (c *ClassStudent) DeleteStudentFromClass(ctx context.Context, classID, studentID string) error {
	return c.classStudentRepo.MarkLeft(ctx, classID, studentID)
}
