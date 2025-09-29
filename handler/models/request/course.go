package request

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
	StudentIDs []string `json:"student_ids"`
	CourseID   string   `json:"course_id"`
}

func (r *EnrollCourseRequest) Validate() error {
	if len(r.StudentIDs) == 0 {
		return errors.New("student_ids is required")
	}

	if r.CourseID == "" {
		return errors.New("course_id is required")
	}

	return nil
}

type UpdateCourseRequest struct {
	Fields map[string]interface{} `json:"fields" binding:"required"`
}

func (r *UpdateCourseRequest) Validate() error {
	if len(r.Fields) == 0 {
		return errors.New("no field provided")
	}

	allowed := map[string]bool{
		"name":              true,
		"description":       true,
		"session_count":     true,
		"price_per_session": true,
	}

	for k, v := range r.Fields {
		if !allowed[k] {
			return errors.New("field is invalid")
		}

		switch k {
		case "name":
			name, ok := v.(string)
			if !ok || len(name) < 2 {
				return errors.New("invalid name: must be a string with length >= 2")
			}
		case "session_count":
			sc, ok := v.(int)
			if !ok || sc < 0 {
				return errors.New("invalid session_count: must be a positive number")
			}
		case "price_per_session":
			pps, ok := v.(int)
			if !ok || pps < 0 {
				return errors.New("invalid age: must be a positive number")
			}
		}
	}

	return nil
}

type SearchCoursesRequest struct {
	Name         *string `form:"name"`
	Description  *string `form:"description"`
	SessionCount *int    `form:"session_count"`
	PriceMin     *int    `form:"price_per_session"`
	PriceMax     *int    `form:"price_per_session"`
	OrderBy      *string `form:"order_by"`
}
