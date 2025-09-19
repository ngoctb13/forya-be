package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ICourseStudentRepo interface {
	Create(ctx context.Context, cs *models.CourseStudent) error
}
