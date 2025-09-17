package models

type CreateClassRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EnrollStudentRequest struct {
	ClassID   string `json:"class_id"`
	StudentID string `json:"student_id"`
}
