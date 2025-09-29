package models

import "time"

type CourseStudent struct {
	StudentID     string    `json:"student_id"`
	CourseID      string    `json:"course_id"`
	RemainSession int       `json:"remain_session"`
	IsCompleted   bool      `json:"is_completed"`
	CompletedAt   time.Time `json:"completed_at"`
	StartedAt     time.Time `json:"started_at"`
}
