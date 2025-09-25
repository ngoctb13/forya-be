package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var phoneRegex = regexp.MustCompile(`^[0-9]{9,15}$`)

type Student struct {
	ID                string `json:"id"`
	FullName          string `json:"full_name"`
	Age               int    `json:"age"`
	PhoneNumber       string `json:"phone_number"`
	ParentPhoneNumber string `json:"parent_phone_number"`
	Note              string `json:"note"`
	IsActive          bool   `json:"is_active"`
}

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
			age, ok := v.(int)
			if !ok || age <= 0 {
				return errors.New("invalid age: must be a positive number")
			}
		case "phone_number", "parent_phone_number":
			phone, ok := v.(string)
			if !ok || !phoneRegex.MatchString(phone) {
				return fmt.Errorf("invalid %s: must be a valid phone number", k)
			}
		case "note":
			_, ok := v.(string)
			if !ok {
				return errors.New("invalid note: must be a string")
			}
		}
	}

	return nil
}

type SearchStudentsRequest struct {
	FullName          *string `form:"full_name"`
	AgeMin            *int    `form:"age_min"`
	AgeMax            *int    `form:"age_max"`
	PhoneNumber       *string `form:"phone_number"`
	ParentPhoneNumber *string `form:"parent_phone_number"`
}
