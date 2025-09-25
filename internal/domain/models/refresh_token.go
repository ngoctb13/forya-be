package models

import "time"

type RefreshToken struct {
	ID        string
	Token     string
	UserID    string
	Role      string
	ExpiresAt time.Time
	Revoked   bool
}
