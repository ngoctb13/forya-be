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

func (s *Student) ListClassStudents(ctx context.Context, input *inputs.ListClassStudentsInput) ([]*models.ClassEnrollments, error) {
	opts := models.QueryOptions{
		JoinedAt: input.JoinedAt,
		LeftAt:   input.LeftAt,
	}

	return s.studentRepo.GetStudentsByClassID(ctx, input.ClassID, opts)
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

func (s *Student) ListStudents(ctx context.Context, input *inputs.ListStudentsInput) ([]*models.Student, error) {
	return s.studentRepo.List(ctx, &models.ListFilter{
		FullName:          input.FullName,
		AgeMin:            input.AgeMin,
		AgeMax:            input.AgeMax,
		PhoneNumber:       input.PhoneNumber,
		ParentPhoneNumber: input.ParentPhoneNumber,
	})
}
