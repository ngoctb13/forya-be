package usecases

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/course/repos"
)

type Course struct {
	courseRepo repos.ICourseRepo
}

func NewCourse(courseRepo repos.ICourseRepo) *Course {
	return &Course{
		courseRepo: courseRepo,
	}
}

func (c *Course) CreateCourse(ctx context.Context, input *models.CreateCourseInput) error {
	course := &models.Course{
		Name:            input.Name,
		Description:     input.Description,
		SessionCount:    input.SessionCount,
		PricePerSession: input.PricePerSession,
		IsActive:        true,
	}

	return c.courseRepo.Create(ctx, course)
}
