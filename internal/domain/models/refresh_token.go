package models

import "time"

type RefreshToken struct {
	Token     string
	UserID    string
	Role      string
	ExpiresAt time.Time
	Revoked   bool
}
