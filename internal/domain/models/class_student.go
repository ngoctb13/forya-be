package models

import "time"

type ClassStudent struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	ClassID   string    `json:"class_id"`
	StudentID string    `json:"student_id"`
	JoinedAt  time.Time `json:"joined_at"`
	LeftAt    time.Time `json:"left_at"`
}

type EnrollStudentInput struct {
	ClassID   string `json:"class_id"`
	StudentID string `json:"student_id"`
}
