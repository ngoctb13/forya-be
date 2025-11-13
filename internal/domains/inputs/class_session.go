package inputs

import "time"

type CreateClassSessionInput struct {
	Name    string
	ClassID string
	HeldAt  time.Time
}

type ListClassSessionsInput struct {
	ClassID   *string
	StartTime *time.Time
	EndTime   *time.Time
	Page      int
	Limit     int
}

type MarkClassSessionAttendanceInput struct {
	ClassSessionID  string
	CourseStudentID string
	IsAttended      bool
}
