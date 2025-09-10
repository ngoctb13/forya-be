package repos

import (
	classRp "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	userRp "github.com/ngoctb13/forya-be/internal/domains/user/repos"
)

type IRepo interface {
	Users() userRp.IUserRepo
	Classes() classRp.IClassRepo
}
