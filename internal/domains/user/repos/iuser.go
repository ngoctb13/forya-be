package repos

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
}
