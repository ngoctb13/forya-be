package response

import (
	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/outputs"
)

type ListClassSessionsResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	HeldAt string `json:"held_at"`
	Class  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"class"`
	Pagination
}

func ToListClassSessionsResponse(input []*outputs.ListClassSessionsOutput, inPagination *models.Pagination) []*ListClassSessionsResponse {
	var res []*ListClassSessionsResponse
	for _, v := range input {
		in := &ListClassSessionsResponse{
			ID:     v.ID,
			Name:   v.Name,
			HeldAt: v.HeldAt.String(),
		}
		in.Class.ID = v.Class.ID
		in.Class.Name = v.Class.Name
		in.Pagination = ToPagination(inPagination)

		res = append(res, in)
	}

	return res
}
