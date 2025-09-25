package repos

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type refreshTokenSQLRepo struct {
	db *gorm.DB
}

func NewRefreshTokenSQLRepo(db *gorm.DB) *refreshTokenSQLRepo {
	return &refreshTokenSQLRepo{
		db: db,
	}
}

func (r *refreshTokenSQLRepo) Create(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *refreshTokenSQLRepo) GetByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	err := r.db.WithContext(ctx).First(&rt, "token = ?", token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rt, nil
}

func (r *refreshTokenSQLRepo) Revoke(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("token = ?", token).
		Update("revoked", true).Error
}
