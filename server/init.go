package server

import (
	"github.com/gin-contrib/cors"
	"github.com/ngoctb13/forya-be/handler"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Domains struct {
	User  *userUC.User
	Class *classUC.Class
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

func (s *Server) InitRouter(domains *Domains) {
	hdl := handler.NewHandler(domains.User, domains.Class)

	authRouter := s.router.Group("api/auth")
	hdl.ConfigRouteAuth(authRouter)

	apiRouter := s.router.Group("api/v1")
	apiRouter.Use(middlewares.AuthMiddleware())
	hdl.ConfigRouteAPI(apiRouter)
}
