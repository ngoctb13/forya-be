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

func (u *userSQLRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).First(&user, "id = ?", id).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).First(&user, "user_name = ?", username).Error
	return &user, err
}
