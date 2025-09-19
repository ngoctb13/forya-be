package models

import "errors"

type CreateCourseRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	SessionCount    int    `json:"session_count"`
	PricePerSession int    `json:"price_per_session"`
}

func (r *CreateCourseRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.SessionCount <= 0 {
		return errors.New("session_count is required")
	}

	if r.PricePerSession <= 0 {
		return errors.New("price_per_session is required")
	}

	return nil
}

type EnrollCourseRequest struct {
	StudentID string `json:"student_id"`
	CourseID  string `json:"course_id"`
}

func (r *EnrollCourseRequest) Validate() error {
	if r.StudentID == "" {
		return errors.New("student_id is required")
	}

	if r.CourseID == "" {
		return errors.New("course_id is required")
	}

	return nil
}
