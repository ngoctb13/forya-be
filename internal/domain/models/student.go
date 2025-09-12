package models

import "time"

type Student struct {
	ID                string `json:"id" gorm:"default:uuid_generate_v4()"`
	FullName          string `json:"full_name"`
	Age               int    `json:"age"`
	PhoneNumber       string `json:"phone_number"`
	ParentPhoneNumber string `json:"parent_phone_number"`
	Note              string `json:"note"`
	IsActive          bool   `json:"is_active"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
