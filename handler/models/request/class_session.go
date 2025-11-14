package request

import (
	"errors"
	"fmt"
)

type CreateClassSessionRequest struct {
	Name    string `json:"name"`
	ClassID string `json:"class_id"`
	HeldAt  string `json:"held_at"`
}

type ListClassSessionsRequest struct {
	ClassID   *string `form:"class_id"`
	StartTime *string `form:"start_time"`
	EndTime   *string `form:"end_time"`
	Page      int     `form:"page"`
	Limit     int     `form:"limit"`
}

type AttendanceItemRequest struct {
	CourseStudentID string `json:"course_student_id"`
	IsAttended      bool   `json:"is_attended"`
}

type BatchMarkClassSessionAttendanceRequest struct {
	Attendances []AttendanceItemRequest `json:"attendances" binding:"required"`
}

func (r *BatchMarkClassSessionAttendanceRequest) Validate() error {
	if len(r.Attendances) == 0 {
		return errors.New("attendances list cannot be empty")
	}
	for i, att := range r.Attendances {
		if att.CourseStudentID == "" {
			return fmt.Errorf("course_student_id is required for attendance item at index %d", i)
		}
	}
	return nil
}
