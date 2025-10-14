package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/internal/domains/student/repos"
)

type Student struct {
	studentRepo repos.IStudentRepo
}

func NewStudent(studentRepo repos.IStudentRepo) *Student {
	return &Student{
		studentRepo: studentRepo,
	}
}

func (s *Student) CreateStudent(ctx context.Context, input *inputs.CreateStudentInput) error {
	student := &models.Student{
		FullName:          input.FullName,
		Age:               input.Age,
		PhoneNumber:       input.PhoneNumber,
		ParentPhoneNumber: input.ParentPhoneNumber,
		Note:              input.Note,
		IsActive:          true,
	}

	return s.studentRepo.CreateStudent(ctx, student)
}

func (s *Student) CreateStudents(ctx context.Context, inputs []*inputs.CreateStudentInput) error {
	var studentArr []*models.Student

	for _, input := range inputs {
		student := &models.Student{
			FullName:          input.FullName,
			Age:               input.Age,
			PhoneNumber:       input.PhoneNumber,
			ParentPhoneNumber: input.ParentPhoneNumber,
			Note:              input.Note,
			IsActive:          true,
		}

		studentArr = append(studentArr, student)
	}

	return s.studentRepo.BatchCreate(ctx, studentArr)
}

func (s *Student) ListClassStudents(ctx context.Context, input *inputs.ListClassStudentsInput) ([]*models.ClassEnrollments, *models.Pagination, error) {
	queries := make(map[string]interface{})
	if input.JoinedAt != nil {
		queries["joined_at"] = input.JoinedAt
	}
	if input.LeftAt != nil {
		queries["left_at"] = input.LeftAt
	}

	pagination := models.NewPagination(input.Page, input.Limit)

	return s.studentRepo.GetStudentsByClassID(ctx, input.ClassID, queries, pagination)
}

func (s *Student) UpdateStudent(ctx context.Context, input *inputs.UpdateStudentInput) (*models.Student, error) {
	student, err := s.studentRepo.GetStudentByID(ctx, input.StudentID)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("student not found")
	}

	return s.studentRepo.UpdateWithMap(ctx, input.StudentID, input.Fields)
}

func (s *Student) ListStudents(ctx context.Context, input *inputs.ListStudentsInput) ([]*models.Student, *models.Pagination, error) {
	queries := make(map[string]interface{})
	if input.FullName != nil {
		queries["name"] = *input.FullName
	}
	if input.AgeMin != nil {
		queries["age_min"] = *input.AgeMin
	}
	if input.AgeMax != nil {
		queries["age_max"] = *input.AgeMax
	}
	if input.PhoneNumber != nil {
		queries["phone_number"] = *input.PhoneNumber
	}
	if input.ParentPhoneNumber != nil {
		queries["parent_phone_number"] = *input.ParentPhoneNumber
	}

	pagination := models.NewPagination(input.Page, input.Limit)

	return s.studentRepo.List(ctx, queries, pagination)
}
