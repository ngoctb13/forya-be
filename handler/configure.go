package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	studentUC "github.com/ngoctb13/forya-be/internal/domains/student/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Handler struct {
	user    *userUC.User
	class   *classUC.Class
	student *studentUC.Student
}

func NewHandler(user *userUC.User, class *classUC.Class, student *studentUC.Student) *Handler {
	return &Handler{
		user:    user,
		class:   class,
		student: student,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	router.GET("/hello", h.Hello())
	router.POST("/class/create", middlewares.AdminOnly(), h.CreateClass())
	router.GET("/class/search", h.SearchClassByName())

	router.POST("/student/create", middlewares.AdminOnly(), h.CreateStudent())
}

func (h *Handler) ConfigRouteAuth(router *gin.RouterGroup) {
	router.POST("/login", h.Login())
	router.POST("/register", h.Register())
}
