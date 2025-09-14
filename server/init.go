package server

import (
	"github.com/gin-contrib/cors"
	"github.com/ngoctb13/forya-be/handler"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	"github.com/ngoctb13/forya-be/infra/repos"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	studentUC "github.com/ngoctb13/forya-be/internal/domains/student/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Domains struct {
	User    *userUC.User
	Class   *classUC.Class
	Student *studentUC.Student
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

func (s *Server) InitDomains(repo repos.IRepo) *Domains {
	user := userUC.NewUser(repo.Users())
	class := classUC.NewClass(repo.Classes())
	student := studentUC.NewStudent(repo.Students())
	return &Domains{
		User:    user,
		Class:   class,
		Student: student,
	}
}

func (s *Server) InitRouter(domains *Domains) {
	hdl := handler.NewHandler(domains.User, domains.Class, domains.Student)

	authRouter := s.router.Group("api/auth")
	hdl.ConfigRouteAuth(authRouter)

	apiRouter := s.router.Group("api/v1")
	apiRouter.Use(middlewares.AuthMiddleware())
	hdl.ConfigRouteAPI(apiRouter)
}
