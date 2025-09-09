package models

import (
	"errors"
	"regexp"
	"strings"
)

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateUserRequest) Validate() error {
	if strings.TrimSpace(r.UserName) == "" {
		return errors.New("user name is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(r.Email) {
		return errors.New("invalid email format")
	}

	if len(r.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	if strings.TrimSpace(r.UserName) == "" {
		return errors.New("user name is required")
	}

	if len(r.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
