package response

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/outputs"
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

func ToListCoursesResponse(in *outputs.ListCoursesOutput, inPagination *models.Pagination) *ListCoursesResponse {
	var courses []Course

	for _, v := range in.Courses {
		item := Course{
			ID:              v.ID,
			Name:            v.Name,
			Description:     v.Description,
			SessionCount:    v.SessionCount,
			PricePerSession: v.PricePerSession,
			IsActive:        v.IsActive,
		}
		courses = append(courses, item)
	}

	return &ListCoursesResponse{
		Courses:    courses,
		Pagination: ToPagination(inPagination),
	}
}
