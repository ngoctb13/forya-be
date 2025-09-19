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

func (c *CourseStudent) CreateCourseStudents(ctx context.Context, input *models.CreateCourseStudentsInput) error {
	var csArr []*models.CourseStudent

	course, err := c.courseRepo.GetByID(ctx, input.CourseID)
	if err != nil {
		return err
	}

	if !course.IsActive {
		return ErrCourseNotActive
	}

	for _, id := range input.StudentIDs {
		cs := &models.CourseStudent{
			StudentID:     id,
			CourseID:      input.CourseID,
			RemainSession: course.SessionCount,
			IsCompleted:   false,
		}

		csArr = append(csArr, cs)
	}

	return c.courseStudentRepo.BatchCreate(ctx, csArr)
}
