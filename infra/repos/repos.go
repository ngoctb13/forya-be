package repos

import (
	"gorm.io/gorm"

	"github.com/ngoctb13/forya-be/config"
	classRp "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	studentRp "github.com/ngoctb13/forya-be/internal/domains/student/repos"
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

func (r *Repo) Students() studentRp.IStudentRepo {
	return NewStudentSQLRepo(r.db)
}
