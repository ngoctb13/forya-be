package outputs

import "github.com/ngoctb13/forya-be/internal/domain/models"

type Supply struct {
	ID           string
	Name         string
	Description  string
	Unit         string
	MinThreshold int
	IsActive     bool
}
type ListSuppliesOutput struct {
	Supplies []Supply
}

func ToListSuppliesOutput(in []*models.Supply) *ListSuppliesOutput {
	var supplyArr []Supply
	for _, v := range in {
		item := Supply{
			ID:           v.ID,
			Name:         v.Name,
			Description:  v.Description,
			Unit:         v.Unit,
			MinThreshold: v.MinThreshold,
			IsActive:     v.IsActive,
		}
		supplyArr = append(supplyArr, item)
	}

	return &ListSuppliesOutput{
		Supplies: supplyArr,
	}
}
