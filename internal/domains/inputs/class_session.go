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

type AttendanceItem struct {
	CourseStudentID string
	IsAttended      bool
	Supplies        []SupplyUsageItem
}

type SupplyUsageItem struct {
	SupplyID  string
	Quantity  int
	UnitPrice *float64
}

type BatchMarkClassSessionAttendanceInput struct {
	ClassSessionID string
	Attendances    []AttendanceItem
}
