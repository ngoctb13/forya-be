package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassStudentRepo interface {
	Create(ctx context.Context, cs *models.ClassStudent) error
	MarkLeft(ctx context.Context, classID, studentID string) error
}
