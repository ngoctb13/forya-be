package response

import (
	"time"

	"github.com/ngoctb13/forya-be/internal/domain/models"
)

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

type ClassStudentsResponse struct {
	Student  Student
	JoinedAt time.Time `json:"joined_at"`
	LeftAt   time.Time `json:"left_at"`
}

type ListClassStudentsResponse struct {
	ClassStudents []ClassStudentsResponse `json:"students"`
	Pagination
}

func ToListClassStudentsResponse(inArr []*models.ClassEnrollments, inPagination *models.Pagination) ListClassStudentsResponse {
	var csArr []ClassStudentsResponse

	for _, v := range inArr {
		cs := ClassStudentsResponse{
			Student: Student{
				ID:                v.ID,
				FullName:          v.FullName,
				Age:               v.Age,
				PhoneNumber:       v.PhoneNumber,
				ParentPhoneNumber: v.ParentPhoneNumber,
				Note:              v.Note,
				IsActive:          v.IsActive,
			},
			JoinedAt: v.JoinedAt,
			LeftAt:   *v.LeftAt,
		}

		csArr = append(csArr, cs)
	}

	pagination := Pagination{
		Page:      inPagination.Page,
		Limit:     inPagination.Limit,
		Total:     int(inPagination.Total),
		TotalPage: inPagination.TotalPage,
	}

	return ListClassStudentsResponse{
		ClassStudents: csArr,
		Pagination:    pagination,
	}
}
