package response

import "github.com/ngoctb13/forya-be/internal/domain/models"

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

func ToUserResponse(in *models.User) User {
	return User{
		ID:       in.ID,
		Email:    in.Email,
		UserName: in.UserName,
		Role:     in.Role,
	}
}
