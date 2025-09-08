package repos

import (
	"gorm.io/gorm"

	"github.com/ngoctb13/forya-be/config"
	"github.com/ngoctb13/forya-be/internal/domains/user/repos"
)

type Repo struct {
	db  *gorm.DB
	cfg *config.PostgresConfig
}

func NewSQLRepo(db *gorm.DB, cfg *config.PostgresConfig) IRepo {
	return &Repo{
		db:  db,
		cfg: cfg,
	}
}

func (r *Repo) Users() repos.IUserRepo {
	return NewUserSQLRepo(r.db)
}
