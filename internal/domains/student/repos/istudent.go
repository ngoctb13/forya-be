package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IStudentRepo interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	BatchCreate(ctx context.Context, students []*models.Student) error
	UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error)
	DeleteStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetStudentsByClassID(ctx context.Context, classID string, queryOpts models.QueryOptions) ([]*models.ClassEnrollments, error)
}
