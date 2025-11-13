package response

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type ClassSessionItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	HeldAt string `json:"held_at"`
	Class  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"class"`
}

type ListClassSessionsResponse struct {
	Sessions []ClassSessionItem `json:"sessions"`
	Pagination
}

// ToListClassSessionsResponse maps domain models directly to response (removed outputs layer)
func ToListClassSessionsResponse(sessions []*models.ClassSession, pagination *models.Pagination) *ListClassSessionsResponse {
	var responseSessions []ClassSessionItem

	for _, v := range sessions {
		item := ClassSessionItem{
			ID:     v.ID,
			Name:   v.Name,
			HeldAt: v.HeldAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		if v.Class != nil {
			item.Class.ID = v.Class.ID
			item.Class.Name = v.Class.Name
		}
		responseSessions = append(responseSessions, item)
	}

	return &ListClassSessionsResponse{
		Sessions:   responseSessions,
		Pagination: ToPagination(pagination),
	}
}

type ClassSessionAttendance struct {
	ID              string `json:"id"`
	ClassSessionID  string `json:"class_session_id"`
	CourseStudentID string `json:"course_student_id"`
	IsAttended      bool   `json:"is_attended"`
}

func ToClassSessionAttendance(attendance *models.ClassSessionAttendance) ClassSessionAttendance {
	return ClassSessionAttendance{
		ID:              attendance.ID,
		ClassSessionID:  attendance.ClassSessionID,
		CourseStudentID: attendance.CourseStudentID,
		IsAttended:      attendance.IsAttended,
	}
}
