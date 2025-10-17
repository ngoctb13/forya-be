package server

import (
	"github.com/gin-contrib/cors"
	"github.com/ngoctb13/forya-be/handler"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	"github.com/ngoctb13/forya-be/infra/repos"
	"github.com/ngoctb13/forya-be/infra/txn"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	classSessionUC "github.com/ngoctb13/forya-be/internal/domains/class_session/usecases"
	classStudentUC "github.com/ngoctb13/forya-be/internal/domains/class_student/usecases"
	courseUC "github.com/ngoctb13/forya-be/internal/domains/course/usecases"
	courseStudentUC "github.com/ngoctb13/forya-be/internal/domains/course_student/usecases"
	authUC "github.com/ngoctb13/forya-be/internal/domains/refresh_token/usecases"
	studentUC "github.com/ngoctb13/forya-be/internal/domains/student/usecases"
	supplyUC "github.com/ngoctb13/forya-be/internal/domains/supply/usecases"
	supplyBatchUC "github.com/ngoctb13/forya-be/internal/domains/supply_batch/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Domains struct {
	User          *userUC.User
	Class         *classUC.Class
	Student       *studentUC.Student
	ClassStudent  *classStudentUC.ClassStudent
	Course        *courseUC.Course
	CourseStudent *courseStudentUC.CourseStudent
	Auth          *authUC.Auth
	ClassSession  *classSessionUC.ClassSession
	Supply        *supplyUC.Supply
	SupplyBatch   *supplyBatchUC.SupplyBatch
}

func (s *Server) InitCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
		"X-Google-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) InitDomains(repos repos.IRepo, t txn.ITxn) *Domains {
	user := userUC.NewUser(repos.Users())
	class := classUC.NewClass(repos.Classes())
	student := studentUC.NewStudent(repos.Students())
	classStudent := classStudentUC.NewClassStudent(repos.ClassStudent())
	course := courseUC.NewCourse(repos.Courses())
	courseStudent := courseStudentUC.NewCourseStudent(repos.CourseStudent(), repos.Courses())
	auth := authUC.NewAuth(repos.RefreshToken())
	classSession := classSessionUC.NewClassSession(repos.ClassSession())
	supply := supplyUC.NewSupply(repos.Supply())
	supplyBatch := supplyBatchUC.NewSupply(repos.SupplyBatch(), repos.Supply())
	return &Domains{
		User:          user,
		Class:         class,
		Student:       student,
		ClassStudent:  classStudent,
		Course:        course,
		CourseStudent: courseStudent,
		Auth:          auth,
		ClassSession:  classSession,
		Supply:        supply,
		SupplyBatch:   supplyBatch,
	}
}

func (s *Server) InitRouter(domains *Domains) {
	hdl := handler.NewHandler(domains.User, domains.Class, domains.Student, domains.ClassStudent, domains.Course,
		domains.CourseStudent, domains.Auth, domains.ClassSession, domains.Supply, domains.SupplyBatch)

	authRouter := s.router.Group("api/auth")
	hdl.ConfigRouteAuth(authRouter)

	apiRouter := s.router.Group("api/v1")
	apiRouter.Use(middlewares.AuthMiddleware())
	hdl.ConfigRouteAPI(apiRouter)
}
