package usecases

import "github.com/ngoctb13/forya-be/internal/domains/user/repos"

type User struct {
	userRepo repos.IUserRepo
}

func NewUser(userRepo repos.IUserRepo) *User {
	return &User{
		userRepo: userRepo,
	}
}
