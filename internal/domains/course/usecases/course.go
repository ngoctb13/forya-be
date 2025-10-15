package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/course/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/internal/domains/outputs"
	"github.com/shopspring/decimal"
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
		PricePerSession: decimal.NewFromFloat(input.PricePerSession),
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

func (c *Course) ListCourses(ctx context.Context, input *inputs.ListCoursesInput) (*outputs.ListCoursesOutput, *models.Pagination, error) {
	pagination := models.NewPagination(input.Page, input.Limit)
	queries := make(map[string]interface{})
	if input.Name != nil {
		queries["name"] = input.Name
	}
	if input.SessionCount != nil {
		queries["session_count"] = input.SessionCount
	}
	if input.PriceMin != nil {
		queries["price_min"] = input.PriceMin
	}
	if input.PriceMax != nil {
		queries["price_max"] = input.PriceMax
	}
	if input.OrderBy != nil {
		queries["order_by"] = input.OrderBy
	}

	courseArr, p, err := c.courseRepo.List(ctx, queries, pagination)

	return outputs.ToListCoursesOutput(courseArr), p, err
}
