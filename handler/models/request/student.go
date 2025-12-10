package request

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

var phoneRegex = regexp.MustCompile(`^[0-9]{9,15}$`)

type CreateStudentRequest struct {
	FullName          string `json:"full_name"`
	Age               int    `json:"age"`
	PhoneNumber       string `json:"phone_number"`
	ParentPhoneNumber string `json:"parent_phone_number"`
	Note              string `json:"note"`
}

func (r CreateStudentRequest) Validate() error {
	if strings.TrimSpace(r.FullName) == "" {
		return errors.New("full_name is required")
	}

	if r.Age <= 0 {
		return errors.New("age must be greater than 0")
	}

	if r.PhoneNumber != "" && !phoneRegex.MatchString(r.PhoneNumber) {
		return errors.New("phone_number is invalid, must be 9-15 digits")
	}

	if r.ParentPhoneNumber != "" && !phoneRegex.MatchString(r.ParentPhoneNumber) {
		return errors.New("parent_phone_number is invalid, must be 9-15 digits")
	}

	return nil
}

type UpdateStudentRequest struct {
	Fields map[string]interface{} `json:"fields" binding:"required"`
}

func (r *UpdateStudentRequest) Validate() error {
	if len(r.Fields) == 0 {
		return errors.New("no field provided")
	}

	allowed := map[string]bool{
		"full_name":           true,
		"age":                 true,
		"phone_number":        true,
		"parent_phone_number": true,
		"note":                true,
	}

	for k, v := range r.Fields {
		if !allowed[k] {
			return errors.New("field is invalid")
		}

		switch k {
		case "full_name":
			name, ok := v.(string)
			if !ok || len(name) < 2 {
				return errors.New("invalid name: must be a string with length >= 2")
			}
		case "age":
			age, ok := v.(float64)
			if !ok || int(age) <= 0 {
				return errors.New("invalid age: must be a positive number")
			}
		case "phone_number", "parent_phone_number":
			phone, ok := v.(string)
			if !ok || !phoneRegex.MatchString(phone) {
				return fmt.Errorf("invalid %s: must be a valid phone number", k)
			}
		}
	}

	return nil
}

type ListStudentsRequest struct {
	FullName          *string `form:"full_name"`
	AgeMin            *int    `form:"age_min"`
	AgeMax            *int    `form:"age_max"`
	PhoneNumber       *string `form:"phone_number"`
	ParentPhoneNumber *string `form:"parent_phone_number"`
	Page              int     `form:"page"`
	Limit             int     `form:"limit"`
}

type ListClassStudentsRequest struct {
	JoinedAtAfter *string `form:"joined_at_after"`
	LeftAtAfter   *string `form:"left_at_after"`
	Page          int     `form:"page"`
	Limit         int     `form:"limit"`
}

func (r *ListClassStudentsRequest) ValidateAndMap() (*inputs.ListClassStudentsInput, error) {
	var input inputs.ListClassStudentsInput
	if r.JoinedAtAfter != nil {
		t, err := time.Parse(time.RFC3339, *r.JoinedAtAfter)
		if err != nil {
			return nil, err
		}

		input.JoinedAt = &t
	}

	if r.LeftAtAfter != nil {
		t, err := time.Parse(time.RFC3339, *r.LeftAtAfter)
		if err != nil {
			return nil, err
		}

		input.LeftAt = &t
	}

	input.Page = r.Page
	input.Limit = r.Limit

	return &input, nil
}
