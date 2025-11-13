package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassSessionAttendance interface {
	MarkAttendance(ctx context.Context, sessionID, courseStudentID string, isAttended bool) (*models.ClassSessionAttendance, error)
}
