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

type SupplyPurchaseRequest struct {
	SupplyID  string   `json:"supply_id"`
	Quantity  int      `json:"quantity"`
	UnitPrice *float64 `json:"unit_price"`
}

type AttendanceItemRequest struct {
	CourseStudentID string                  `json:"course_student_id"`
	IsAttended      bool                    `json:"is_attended"`
	Supplies        []SupplyPurchaseRequest `json:"supplies"`
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
		for j, sup := range att.Supplies {
			if sup.SupplyID == "" {
				return fmt.Errorf("supply_id is required for attendance[%d] supply[%d]", i, j)
			}
			if sup.Quantity <= 0 {
				return fmt.Errorf("quantity must be > 0 for attendance[%d] supply[%d]", i, j)
			}
			if sup.UnitPrice != nil && *sup.UnitPrice < 0 {
				return fmt.Errorf("unit_price must be >= 0 for attendance[%d] supply[%d]", i, j)
			}
		}
	}

	return nil
}
