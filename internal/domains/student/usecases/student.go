package usecases

import (
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
