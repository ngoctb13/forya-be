package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"gorm.io/gorm"
)

type userSQLRepo struct {
	db *gorm.DB
}

func NewUserSQLRepo(db *gorm.DB) *userSQLRepo {
	return &userSQLRepo{
		db: db,
	}
}

func (u *userSQLRepo) CreateUser(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}
