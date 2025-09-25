package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IRefreshTokenRepo interface {
	Create(ctx context.Context, rt *models.RefreshToken) error
	GetByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	Revoke(ctx context.Context, token string) error
}
