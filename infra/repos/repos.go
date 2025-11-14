package repos

import (
	"gorm.io/gorm"

	"github.com/ngoctb13/forya-be/config"
	classRp "github.com/ngoctb13/forya-be/internal/domains/class/repos"
	classSessionRp "github.com/ngoctb13/forya-be/internal/domains/class_session/repos"
	classStudentRp "github.com/ngoctb13/forya-be/internal/domains/class_student/repos"
	courseRp "github.com/ngoctb13/forya-be/internal/domains/course/repos"
	courseStudentRp "github.com/ngoctb13/forya-be/internal/domains/course_student/repos"
	refreshTokenRp "github.com/ngoctb13/forya-be/internal/domains/refresh_token/repos"
	studentRp "github.com/ngoctb13/forya-be/internal/domains/student/repos"
	supplyRp "github.com/ngoctb13/forya-be/internal/domains/supply/repos"
	supplyBatchRp "github.com/ngoctb13/forya-be/internal/domains/supply_batch/repos"
	supplyUsageRp "github.com/ngoctb13/forya-be/internal/domains/supply_usage/repos"
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

func (r *Repo) ClassStudent() classStudentRp.IClassStudentRepo {
	return NewClassStudentSQLRepo(r.db)
}

func (r *Repo) Courses() courseRp.ICourseRepo {
	return NewCourseSQLRepo(r.db)
}

func (r *Repo) CourseStudent() courseStudentRp.ICourseStudentRepo {
	return NewCourseStudentSQLRepo(r.db)
}

func (r *Repo) RefreshToken() refreshTokenRp.IRefreshTokenRepo {
	return NewRefreshTokenSQLRepo(r.db)
}

func (r *Repo) ClassSession() classSessionRp.IClassSession {
	return NewClassSessionSQLRepo(r.db)
}

func (r *Repo) ClassSessionAttendance() classSessionRp.IClassSessionAttendance {
	return NewClassSessionAttendanceSQLRepo(r.db)
}

func (r *Repo) Supply() supplyRp.ISupply {
	return NewSupplySQLRepo(r.db)
}

func (r *Repo) SupplyBatch() supplyBatchRp.ISupplyBatch {
	return NewSupplyBatchSQLRepo(r.db)
}

func (r *Repo) SupplyUsage() supplyUsageRp.ISupplyUsage {
	return NewSupplyUsageSQLRepo(r.db)
}
