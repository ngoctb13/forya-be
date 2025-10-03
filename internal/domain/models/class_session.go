package models

import "time"

type ClassSession struct {
	ID      string `gorm:"default:uuid_generate_v4()"`
	ClassID string
	HeldAt  time.Time
}
