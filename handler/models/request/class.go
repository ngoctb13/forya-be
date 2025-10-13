package request

import "errors"

type CreateClassRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateClassRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

type EnrollClassRequest struct {
	ClassID    string   `json:"class_id"`
	StudentIDs []string `json:"student_ids"`
}

func (r *EnrollClassRequest) Validate() error {
	if r.ClassID == "" {
		return errors.New("class_id is required")
	}

	if len(r.StudentIDs) == 0 {
		return errors.New("student_ids is required")
	}

	return nil
}

type ListClassRequest struct {
	Name  *string `form:"name"`
	Page  int     `form:"page"`
	Limit int     `form:"limit"`
}
