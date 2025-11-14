package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassSessionAttendance interface {
	BatchMarkAttendance(ctx context.Context, sessionID string, attendances []*models.ClassSessionAttendance) error
}
