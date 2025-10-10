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

type SearchClassRequest struct {
	Name  *string `form:"name"`
	Page  int     `form:"page"`
	Limit int     `form:"limit"`
}

func (r *SearchClassRequest) Validate() error {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.Limit <= 0 {
		r.Limit = 10
	}

	return nil
}
