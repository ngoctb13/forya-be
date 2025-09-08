package models

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"passwords"`
}
