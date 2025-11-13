package request

import "errors"

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

type MarkClassSessionAttendanceRequest struct {
	CourseStudentID string `json:"course_student_id"`
	IsAttended      bool   `json:"is_attended"`
}

func (r *MarkClassSessionAttendanceRequest) Validate() error {
	if r.CourseStudentID == "" {
		return errors.New("course_student_id is required")
	}
	return nil
}
