package response

import "github.com/ngoctb13/forya-be/internal/domain/models"

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

func ToPagination(in *models.Pagination) Pagination {
	return Pagination{
		Page:      in.Page,
		Limit:     in.Limit,
		Total:     int(in.Total),
		TotalPage: in.TotalPage,
	}
}
