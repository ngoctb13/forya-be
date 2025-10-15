package outputs

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/shopspring/decimal"
)

type Course struct {
	ID              string
	Name            string
	Description     string
	SessionCount    int
	IsActive        bool
	PricePerSession decimal.Decimal
}
type ListCoursesOutput struct {
	Courses []Course
}

func ToListCoursesOutput(in []*models.Course) *ListCoursesOutput {
	var courseArr []Course
	for _, v := range in {
		c := Course{
			ID:              v.ID,
			Name:            v.Name,
			Description:     v.Description,
			SessionCount:    v.SessionCount,
			IsActive:        v.IsActive,
			PricePerSession: v.PricePerSession,
		}

		courseArr = append(courseArr, c)
	}

	out := &ListCoursesOutput{
		Courses: courseArr,
	}

	return out
}
