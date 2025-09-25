package usecases

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/refresh_token/repos"
	"github.com/ngoctb13/forya-be/pkg/auth"
)

type Auth struct {
	refreshTokenRp repos.IRefreshTokenRepo
}

func NewAuth(refreshTokenRp repos.IRefreshTokenRepo) *Auth {
	return &Auth{
		refreshTokenRp: refreshTokenRp,
	}
}

func (a *Auth) GenerateAccessToken(userID string, role string) (string, error) {
	return auth.GenerateJWT(userID, role)
}

func (a *Auth) GenerateRefreshToken(ctx context.Context, userID string, role string) (*models.RefreshToken, error) {
	token, err := a.generateRandomString(64)
	if err != nil {
		return nil, err
	}

	rt := &models.RefreshToken{
		Token:     token,
		UserID:    userID,
		Role:      role,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Revoked:   false,
	}
	if err := a.refreshTokenRp.Create(ctx, rt); err != nil {
		return nil, err
	}

	return rt, nil
}

func (a *Auth) RefreshAccessToken(ctx context.Context, token string) (string, *models.RefreshToken, error) {
	rt, err := a.refreshTokenRp.GetByToken(ctx, token)
	if err != nil || rt == nil || rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		return "", nil, errors.New("invalid refresh token")
	}

	newRT, err := a.GenerateRefreshToken(ctx, rt.UserID, rt.Role)
	if err != nil {
		return "", nil, err
	}

	err = a.refreshTokenRp.Revoke(ctx, rt.Token)
	if err != nil {
		return "", nil, err
	}

	at, err := a.GenerateAccessToken(rt.UserID, rt.Role)
	if err != nil {
		return "", nil, err
	}

	return at, newRT, nil
}

func (a *Auth) generateRandomString(len int) (string, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
