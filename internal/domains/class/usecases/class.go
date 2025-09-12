package usecases

import (
	"context"
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/class/repos"
)

type Class struct {
	classRepo repos.IClassRepo
}

func NewClass(classRepo repos.IClassRepo) *Class {
	return &Class{
		classRepo: classRepo,
	}
}

func (c *Class) CreateClass(ctx context.Context, input *models.CreateClassInput) error {
	class := &models.Class{
		Name:        input.Name,
		Description: input.Description,
		Schedule:    time.Now().Weekday().String(),
	}

	return c.classRepo.CreateClass(ctx, class)
}

func (c *Class) SearchClassByName(ctx context.Context, name string) ([]*models.Class, error) {
	return c.classRepo.GetClassContainName(ctx, name)
}
