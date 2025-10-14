package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IStudentRepo interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	BatchCreate(ctx context.Context, students []*models.Student) error
	DeleteStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetStudentsByClassID(ctx context.Context, classID string, queries map[string]interface{}, pagination *models.Pagination) ([]*models.ClassEnrollments, *models.Pagination, error)
	UpdateWithMap(ctx context.Context, studentID string, fields map[string]interface{}) (*models.Student, error)
	List(ctx context.Context, queries map[string]interface{}, pagination *models.Pagination) ([]*models.Student, *models.Pagination, error)
}
