package request

import (
	"errors"
	"strings"

	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type CreateCourseRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	SessionCount    int     `json:"session_count"`
	PricePerSession float64 `json:"price_per_session"`
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
	Name            *string  `json:"name"`
	Description     *string  `json:"description"`
	SessionCount    *int     `json:"session_count"`
	PricePerSession *float64 `json:"price_per_session"`
}

func (r *UpdateCourseRequest) Validate() error {
	if r.Name == nil && r.Description == nil && r.SessionCount == nil && r.PricePerSession == nil {
		return errors.New("no field provided")
	}
	if r.Name != nil && len(*r.Name) < 2 {
		return errors.New("invalid name: length must be >= 2")
	}
	if r.SessionCount != nil && *r.SessionCount < 0 {
		return errors.New("invalid session_count: must be a non-negative integer")
	}
	if r.PricePerSession != nil && *r.PricePerSession <= 0 {
		return errors.New("invalid price_per_session: must be > 0")
	}
	return nil
}

type ListCoursesRequest struct {
	Name         *string  `form:"name"`
	SessionCount *int     `form:"session_count"`
	PriceMin     *float64 `form:"price_min"`
	PriceMax     *float64 `form:"price_max"`
	OrderBy      *string  `form:"order_by"`
	Page         int      `form:"page"`
	Limit        int      `form:"limit"`
}

func (r *ListCoursesRequest) ValidateAndMap() (*inputs.ListCoursesInput, error) {
	page := r.Page
	limit := r.Limit

	if limit > 100 {
		limit = 100
	}

	var orderBy *string
	if r.OrderBy != nil {
		ob := strings.ToLower(strings.TrimSpace(*r.OrderBy))
		switch ob {
		case "price_asc", "price_desc", "session_count_asc", "session_count_desc":
			orderBy = &ob
		case "":
		default:
			return nil, errors.New("invalid order_by")
		}
	}

	if r.PriceMin != nil && r.PriceMax != nil && *r.PriceMin > *r.PriceMax {
		return nil, errors.New("price_min cannot be greater than price_max")
	}

	var name *string
	if r.Name != nil {
		trimmed := strings.TrimSpace(*r.Name)
		if trimmed != "" {
			name = &trimmed
		}
	}

	input := &inputs.ListCoursesInput{
		Name:         name,
		SessionCount: r.SessionCount,
		PriceMax:     r.PriceMax,
		PriceMin:     r.PriceMin,
		OrderBy:      orderBy,
		Page:         page,
		Limit:        limit,
	}
	return input, nil
}
