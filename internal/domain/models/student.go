package models

import (
	"time"
)

type Student struct {
	ID                string    `json:"id" gorm:"default:uuid_generate_v4()"`
	FullName          string    `json:"full_name"`
	Age               int       `json:"age"`
	PhoneNumber       string    `json:"phone_number"`
	ParentPhoneNumber string    `json:"parent_phone_number"`
	Note              string    `json:"note"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type ClassEnrollments struct {
	Student
	JoinedAt time.Time  `json:"joined_at"`
	LeftAt   *time.Time `json:"left_at"`
}
