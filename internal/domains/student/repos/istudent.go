package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IStudentRepo interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	DeleteStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
}
