package usecases

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	courseRp "github.com/ngoctb13/forya-be/internal/domains/course/repos"
	courseStudentRp "github.com/ngoctb13/forya-be/internal/domains/course_student/repos"
)

type CourseStudent struct {
	courseStudentRepo courseStudentRp.ICourseStudentRepo
	courseRepo        courseRp.ICourseRepo
}

func NewCourseStudent(csr courseStudentRp.ICourseStudentRepo, cr courseRp.ICourseRepo) *CourseStudent {
	return &CourseStudent{
		courseStudentRepo: csr,
		courseRepo:        cr,
	}
}

func (c *CourseStudent) CreateCourseStudent(ctx context.Context, input *models.CreateCourseStudentInput) error {
	course, err := c.courseRepo.GetByID(ctx, input.CourseID)
	if err != nil {
		return err
	}

	if !course.IsActive {
		return ErrCourseNotActive
	}

	cs := &models.CourseStudent{
		StudentID:     input.StudentID,
		CourseID:      input.CourseID,
		RemainSession: course.SessionCount,
		IsCompleted:   false,
	}

	return c.courseStudentRepo.Create(ctx, cs)
}
