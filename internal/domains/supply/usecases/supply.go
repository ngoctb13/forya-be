package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/internal/domains/outputs"
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

func (s *Supply) GetSuppliesByName(ctx context.Context, input string) ([]*outputs.GetSuppliesByNameOutput, error) {
	supArr, err := s.supplyRepo.ListByName(ctx, input)
	if err != nil {
		return nil, err
	}

	if len(supArr) == 0 {
		return nil, errors.New("do no have any supply")
	}

	return outputs.ToGetSuppliesByNameOutput(supArr), nil
}

func (s *Supply) DeleteSupply(ctx context.Context, id string) error {
	return s.supplyRepo.Delete(ctx, id)
}
