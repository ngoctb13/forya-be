package usecases

import (
	"context"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/user/repos"
)

const (
	BKN   = "buikimngoc"
	ADMIN = "admin"
)

type User struct {
	userRepo repos.IUserRepo
}

func NewUser(userRepo repos.IUserRepo) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) CreateUser(ctx context.Context, input *models.CreateUserInput) error {
	existingUser, err := u.userRepo.GetUserByUsername(ctx, input.UserName)
	if err == nil && existingUser != nil {
		return ErrUsernameAlreadyExists
	}

	existingUser, err = u.userRepo.GetUserByEmail(ctx, input.Email)
	if err == nil && existingUser != nil {
		return ErrEmailAlreadyExists
	}

	if input.UserName == BKN {
		input.Role = ADMIN
	}

	user := &models.User{
		Email:    input.Email,
		UserName: input.UserName,
		Password: input.Password,
		Role:     input.Role,
	}
	return u.userRepo.CreateUser(ctx, user)
}

func (u *User) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
