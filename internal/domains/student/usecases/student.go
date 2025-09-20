package usecases

import (
	"context"
	"github.com/ngoctb13/forya-be/internal/domain/models"
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

func (s *Student) CreateStudent(ctx context.Context, input *models.CreateStudentInput) error {
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

func (s *Student) CreateStudents(ctx context.Context, inputs []*models.CreateStudentInput) error {
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
