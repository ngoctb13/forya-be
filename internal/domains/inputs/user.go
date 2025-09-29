package inputs

type CreateUserInput struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
