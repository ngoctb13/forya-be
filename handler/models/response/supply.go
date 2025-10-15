package response

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/outputs"
)

type Supply struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Unit         string `json:"unit"`
	MinThreshold int    `json:"min_threshold"`
	IsActive     bool   `json:"is_active"`
}

func ToSupply(in outputs.Supply) Supply {
	return Supply{
		ID:           in.ID,
		Name:         in.Name,
		Description:  in.Description,
		Unit:         in.Unit,
		MinThreshold: in.MinThreshold,
		IsActive:     in.IsActive,
	}
}

type ListSuppliesResponse struct {
	Supplies []Supply `json:"supplies"`
	Pagination
}

func ToListSuppliesResponse(in *outputs.ListSuppliesOutput, inPagination *models.Pagination) *ListSuppliesResponse {
	var supArr []Supply
	for _, v := range in.Supplies {
		item := ToSupply(v)
		supArr = append(supArr, item)
	}

	return &ListSuppliesResponse{
		Supplies:   supArr,
		Pagination: ToPagination(inPagination),
	}
}
