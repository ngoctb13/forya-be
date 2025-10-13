package usecases

import (
	"context"
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/class/repos"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type Class struct {
	classRepo repos.IClassRepo
}

func NewClass(classRepo repos.IClassRepo) *Class {
	return &Class{
		classRepo: classRepo,
	}
}

func (c *Class) CreateClass(ctx context.Context, input *inputs.CreateClassInput) error {
	class := &models.Class{
		Name:        input.Name,
		Description: input.Description,
		Schedule:    time.Now().Weekday().String(),
		CreatedAt:   time.Now(),
		IsActive:    true,
	}

	return c.classRepo.CreateClass(ctx, class)
}

func (c *Class) SearchClassByName(ctx context.Context, input *inputs.SearchClassByNameInput) ([]*models.Class, *models.Pagination, error) {
	pagination := models.NewPagination(input.Page, input.Limit)
	return c.classRepo.GetClassContainName(ctx, input.Name, pagination)
}
