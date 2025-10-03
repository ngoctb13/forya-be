package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type classSessionSQLRepo struct {
	db *gorm.DB
}

func NewClassSessionSQLRepo(db *gorm.DB) *classSessionSQLRepo {
	return &classSessionSQLRepo{
		db: db,
	}
}

func (r *classSessionSQLRepo) Create(ctx context.Context, session *models.ClassSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}
