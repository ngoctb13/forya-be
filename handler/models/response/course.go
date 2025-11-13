package response

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/shopspring/decimal"
)

type Course struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	SessionCount    int             `json:"session_count"`
	PricePerSession decimal.Decimal `json:"price_per_session"`
	IsActive        bool            `json:"is_active"`
}

type ListCoursesResponse struct {
	Courses []Course `json:"courses"`
	Pagination
}

// ToListCoursesResponse maps domain models directly to response (removed outputs layer)
func ToListCoursesResponse(courses []*models.Course, pagination *models.Pagination) *ListCoursesResponse {
	var responseCourses []Course

	for _, v := range courses {
		item := Course{
			ID:              v.ID,
			Name:            v.Name,
			Description:     v.Description,
			SessionCount:    v.SessionCount,
			PricePerSession: v.PricePerSession,
			IsActive:        v.IsActive,
		}
		responseCourses = append(responseCourses, item)
	}

	return &ListCoursesResponse{
		Courses:    responseCourses,
		Pagination: ToPagination(pagination),
	}
}
