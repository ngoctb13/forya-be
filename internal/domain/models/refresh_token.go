package models

import "time"

type RefreshToken struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Token     string    `json:"token"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	ExpiresAt time.Time `json:"expires_at"`
	Revoked   bool      `json:"revoked"`
}
