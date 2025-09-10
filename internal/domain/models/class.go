package models

import "time"

type Class struct {
	ID          string    `json:"id" gorm:"default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Schedule    string    `json:"schedule"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
}

type CreateClassInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
