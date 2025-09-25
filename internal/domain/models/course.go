package models

import "time"

type Course struct {
	ID              string `json:"id" gorm:"default:uuid_generate_v4()"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	SessionCount    int    `json:"session_count"`
	PricePerSession int    `json:"price_per_session"`
	IsActive        bool   `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CreateCourseInput struct {
	Name            string
	Description     string
	SessionCount    int
	PricePerSession int
}

type UpdateCourseInput struct {
	CourseID string
	Fields   map[string]interface{}
}
