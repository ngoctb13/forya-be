package models

import "time"

type ClassStudent struct {
	ID             string    `json:"id" gorm:"default:uuid_generate_v4()"`
	ClassID        string    `json:"class_id"`
	StudentID      string    `json:"student_id"`
	ParticipatedAt time.Time `json:"participated_at"`
}
