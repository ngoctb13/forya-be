package repos

import (
	"gorm.io/gorm"

	"github.com/ngoctb13/forya-be/config"
	classRp "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	userRp "github.com/ngoctb13/forya-be/internal/domains/user/repos"
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

func (r *Repo) Users() userRp.IUserRepo {
	return NewUserSQLRepo(r.db)
}

func (r *Repo) Classes() classRp.IClassRepo {
	return NewClassSQLRepo(r.db)
}
