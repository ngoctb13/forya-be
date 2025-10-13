package outputs

import "github.com/ngoctb13/forya-be/internal/domain/models"

type GetSuppliesByNameOutput struct {
	ID           string
	Name         string
	Description  string
	Unit         string
	MinThreshold int
}

func ToGetSuppliesByNameOutput(in []*models.Supply) []*GetSuppliesByNameOutput {
	var outs []*GetSuppliesByNameOutput
	for _, s := range in {
		outs = append(outs, &GetSuppliesByNameOutput{
			ID:           s.ID,
			Name:         s.Name,
			Description:  s.Description,
			Unit:         s.Unit,
			MinThreshold: s.MinThreshold,
		})
	}

	return outs
}
