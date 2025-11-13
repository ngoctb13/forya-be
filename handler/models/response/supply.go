package response

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
)

type Supply struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Unit         string `json:"unit"`
	MinThreshold int    `json:"min_threshold"`
	IsActive     bool   `json:"is_active"`
}

type ListSuppliesResponse struct {
	Supplies []Supply `json:"supplies"`
	Pagination
}

// ToListSuppliesResponse maps domain models directly to response (removed outputs layer)
func ToListSuppliesResponse(supplies []*models.Supply, pagination *models.Pagination) *ListSuppliesResponse {
	var responseSupplies []Supply

	for _, v := range supplies {
		item := Supply{
			ID:           v.ID,
			Name:         v.Name,
			Description:  v.Description,
			Unit:         v.Unit,
			MinThreshold: v.MinThreshold,
			IsActive:     v.IsActive,
		}
		responseSupplies = append(responseSupplies, item)
	}

	return &ListSuppliesResponse{
		Supplies:   responseSupplies,
		Pagination: ToPagination(pagination),
	}
}
