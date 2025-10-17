package response

import "github.com/ngoctb13/forya-be/internal/domain/models"

type Class struct {
	ID          string `json:"id" `
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    string `json:"schedule"`
	IsActive    bool   `json:"is_active"`
}

type SearchClassResponse struct {
	Classes []Class `json:"classes"`
	Pagination
}

func ToSearchClassResponse(inArr []*models.Class, inPagination *models.Pagination) SearchClassResponse {
	var arrClass []Class
	for _, v := range inArr {
		arrClass = append(arrClass, Class{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Schedule:    v.Schedule,
			IsActive:    v.IsActive,
		})
	}

	return SearchClassResponse{
		Classes:    arrClass,
		Pagination: ToPagination(inPagination),
	}
}
