package usecases

import "errors"

var (
	ErrCourseNotFound  = errors.New("course not found")
	ErrCourseNotActive = errors.New("course not active")
)
