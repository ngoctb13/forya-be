package response

import "github.com/ngoctb13/forya-be/internal/domain/models"

type Student struct {
	ID                string `json:"id"`
	FullName          string `json:"full_name"`
	Age               int    `json:"age"`
	PhoneNumber       string `json:"phone_number"`
	ParentPhoneNumber string `json:"parent_phone_number"`
	Note              string `json:"note"`
	IsActive          bool   `json:"is_active"`
}

type ListStudentsResponse struct {
	Students []Student `json:"students"`
	Pagination
}

func ToListStudentsResponse(inArr []*models.Student, inPagination *models.Pagination) ListStudentsResponse {
	var students []Student

	for _, v := range inArr {
		s := Student{
			ID:                v.ID,
			FullName:          v.FullName,
			Age:               v.Age,
			PhoneNumber:       v.PhoneNumber,
			ParentPhoneNumber: v.ParentPhoneNumber,
			Note:              v.Note,
			IsActive:          v.IsActive,
		}

		students = append(students, s)
	}

	pagination := Pagination{
		Page:      inPagination.Page,
		Limit:     inPagination.Limit,
		Total:     int(inPagination.Total),
		TotalPage: inPagination.TotalPage,
	}

	return ListStudentsResponse{
		Students:   students,
		Pagination: pagination,
	}
}
