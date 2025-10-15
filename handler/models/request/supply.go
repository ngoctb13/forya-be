package request

import (
	"errors"

	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

type CreateSupplyRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Unit         string `json:"unit"`
	MinThreshold int    `json:"min_threshold"`
}

func (r *CreateSupplyRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}

	if r.Unit == "" {
		return errors.New("unit is required")
	}

	if r.MinThreshold < 0 {
		return errors.New("min_threshold is required")
	}

	return nil
}

type ListSuppliesRequest struct {
	Name         *string `form:"name"`
	MinThreshold *int    `form:"min_threshold"`
	Page         int     `form:"page"`
	Limit        int     `form:"limit"`
}

func (r *ListSuppliesRequest) ValidateAndMap() (*inputs.ListSuppliesInput, error) {
	input := &inputs.ListSuppliesInput{}
	if r.Name != nil {
		input.Name = r.Name
	}

	if r.MinThreshold != nil {
		if *r.MinThreshold < 0 {
			return nil, errors.New("min_threshold can not be negative")
		}
		input.MinThreshold = r.MinThreshold
	}

	input.Page = r.Page
	input.Limit = r.Limit

	return input, nil
}

type UpdateSupplyRequest struct {
	Fields map[string]interface{} `json:"fields" binding:"required"`
}

func (r *UpdateSupplyRequest) Validate() error {
	if len(r.Fields) == 0 {
		return errors.New("no field provided")
	}

	allowed := map[string]bool{
		"name":          true,
		"description":   true,
		"unit":          true,
		"min_threshold": true,
	}

	for k, v := range r.Fields {
		if !allowed[k] {
			return errors.New("field is invalid")
		}

		switch k {
		case "name":
			name, ok := v.(string)
			if !ok || len(name) < 2 {
				return errors.New("invalid name: must be a string with length >= 2")
			}
		case "unit":
			unit, ok := v.(string)
			if !ok || len(unit) == 0 {
				return errors.New("invalid unit")
			}
		case "min_threshold":
			mt, ok := v.(int)
			if !ok || mt < 0 {
				return errors.New("invalid age: must be a positive number")
			}
		}
	}

	return nil
}
