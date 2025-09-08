package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
