package models

import "time"

type ClassStudent struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	ClassID   string    `json:"class_id"`
	StudentID string    `json:"student_id"`
	JoinedAt  time.Time `json:"joined_at"`
	LeftAt    time.Time `json:"left_at"`
}

func (ClassStudent) TableName() string {
	return "class_student"
}

type QueryOptions struct {
	JoinedAt *time.Time
	LeftAt   *time.Time
}

type EnrollClassInput struct {
	ClassID    string
	StudentIDs []string
}
