package inputs

import "time"

type ListClassStudentsInput struct {
	ClassID  string
	JoinedAt *time.Time
	LeftAt   *time.Time
}
type CreateStudentInput struct {
	FullName          string
	Age               int
	PhoneNumber       string
	ParentPhoneNumber string
	Note              string
}

type UpdateStudentInput struct {
	StudentID string
	Fields    map[string]interface{}
}

type ListStudentsInput struct {
	FullName          *string
	AgeMin            *int
	AgeMax            *int
	PhoneNumber       *string
	ParentPhoneNumber *string
}
