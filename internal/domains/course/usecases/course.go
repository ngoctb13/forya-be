package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/course/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type Course struct {
	courseRepo repos.ICourseRepo
}

func NewCourse(courseRepo repos.ICourseRepo) *Course {
	return &Course{
		courseRepo: courseRepo,
	}
}

func (c *Course) CreateCourse(ctx context.Context, input *inputs.CreateCourseInput) error {
	course := &models.Course{
		Name:            input.Name,
		Description:     input.Description,
		SessionCount:    input.SessionCount,
		PricePerSession: input.PricePerSession,
		IsActive:        true,
	}

	return c.courseRepo.Create(ctx, course)
}

func (c *Course) UpdateCourse(ctx context.Context, input *inputs.UpdateCourseInput) (*models.Course, error) {
	course, err := c.courseRepo.GetByID(ctx, input.CourseID)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, errors.New("course not found")
	}
	return c.courseRepo.UpdateWithMap(ctx, input.CourseID, input.Fields)
}

func (c *Course) ListCourses(ctx context.Context, input *inputs.ListCoursesInput) ([]*models.Course, error) {
	return c.courseRepo.GetAll(ctx, &models.GetAllFilter{
		Name:         input.Name,
		Description:  input.Description,
		SessionCount: input.SessionCount,
		PriceMax:     input.PriceMax,
		PriceMin:     input.PriceMin,
		OrderBy:      input.OrderBy,
	})
}
