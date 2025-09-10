package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassRepo interface {
	CreateClass(ctx context.Context, class *models.Class) error
	GetClassByID(ctx context.Context, id string) (*models.Class, error)
	GetClassContainName(ctx context.Context, name string) (*models.Class, error)
}
