package models

import "time"

type Supply struct {
	ID           string `gorm:"default:uuid_generate_v4()"`
	Name         string
	Description  string
	Unit         string
	MinThreshold int
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
