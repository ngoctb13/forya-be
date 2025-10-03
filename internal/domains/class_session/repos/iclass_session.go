package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IClassSession interface {
	Create(ctx context.Context, session *models.ClassSession) error
}
