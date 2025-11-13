package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/course/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
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

	fields := map[string]interface{}{}
	if input.Fields.Name != nil {
		fields["name"] = *input.Fields.Name
	}
	if input.Fields.Description != nil {
		fields["description"] = *input.Fields.Description
	}
	if input.Fields.SessionCount != nil {
		fields["session_count"] = *input.Fields.SessionCount
	}
	if input.Fields.PricePerSession != nil {
		fields["price_per_session"] = decimal.NewFromFloat(*input.Fields.PricePerSession)
	}
	if len(fields) == 0 {
		return course, nil
	}

	return c.courseRepo.UpdateWithMap(ctx, input.CourseID, fields)
}

func (c *Course) ListCourses(ctx context.Context, input *inputs.ListCoursesInput) ([]*models.Course, *models.Pagination, error) {
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

	return courseArr, p, err
}
