package models

import (
	"errors"
	"regexp"
	"strings"
)

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

	phoneRegex := regexp.MustCompile(`^[0-9]{9,15}$`)

	if !phoneRegex.MatchString(r.PhoneNumber) {
		return errors.New("phone_number is invalid, must be 9-15 digits")
	}

	if r.ParentPhoneNumber != "" && !phoneRegex.MatchString(r.ParentPhoneNumber) {
		return errors.New("parent_phone_number is invalid, must be 9-15 digits")
	}

	return nil
}
