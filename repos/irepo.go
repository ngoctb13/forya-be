package repos

import "github.com/ngoctb13/forya-be/internal/domains/user/repos"

type IRepo interface {
	Users() repos.IUserRepo
}
