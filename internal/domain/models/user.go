package models

type User struct {
	ID       string `json:"id" gorm:"default:uuid_generate_v4()"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
