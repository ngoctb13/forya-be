package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassStudentRepo interface {
	Create(ctx context.Context, cs *models.ClassStudent) error
	BatchCreate(ctx context.Context, cs []*models.ClassStudent) error
	MarkLeft(ctx context.Context, classID, studentID string) error
	GetByClassAndStudent(ctx context.Context, classID, studentID string) (*models.ClassStudent, error)
	ListByClassAndStudents(ctx context.Context, classID string, studentIDs []string) ([]*models.ClassStudent, error)
	ResetLeftAtBulk(ctx context.Context, classID string, studentIDs []string) error
}
