package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/internal/domains/supply/repos"
)

type Supply struct {
	supplyRepo repos.ISupply
}

func NewSupply(supplyRepo repos.ISupply) *Supply {
	return &Supply{supplyRepo: supplyRepo}
}

func (s *Supply) CreateSupply(ctx context.Context, input *inputs.CreateSupplyInput) error {
	supply := &models.Supply{
		Name:         input.Name,
		Description:  input.Description,
		Unit:         input.Unit,
		MinThreshold: input.MinThreshold,
		IsActive:     true,
	}

	return s.supplyRepo.Create(ctx, supply)
}

func (s *Supply) UpdateSupply(ctx context.Context, input *inputs.UpdateSupplyInput) error {
	es, err := s.supplyRepo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if es == nil {
		return errors.New("supply not found")
	}

	return s.supplyRepo.UpdateWithFields(ctx, es, input.Fields)
}

// ListSupplies returns domain models directly (removed outputs layer)
func (s *Supply) ListSupplies(ctx context.Context, input *inputs.ListSuppliesInput) ([]*models.Supply, *models.Pagination, error) {
	pagination := models.NewPagination(input.Page, input.Limit)
	queries := make(map[string]interface{})
	if input.Name != nil {
		queries["name"] = input.Name
	}
	if input.MinThreshold != nil {
		queries["min_threshold"] = input.MinThreshold
	}

	supArr, p, err := s.supplyRepo.List(ctx, queries, pagination)

	return supArr, p, err
}

func (s *Supply) DeleteSupply(ctx context.Context, id string) error {
	return s.supplyRepo.Delete(ctx, id)
}
