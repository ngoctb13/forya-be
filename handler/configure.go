package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	classStudentUC "github.com/ngoctb13/forya-be/internal/domains/class_student/usecases"
	studentUC "github.com/ngoctb13/forya-be/internal/domains/student/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Handler struct {
	user         *userUC.User
	class        *classUC.Class
	student      *studentUC.Student
	classStudent *classStudentUC.ClassStudent
}

func NewHandler(user *userUC.User, class *classUC.Class, student *studentUC.Student, classStudent *classStudentUC.ClassStudent) *Handler {
	return &Handler{
		user:         user,
		class:        class,
		student:      student,
		classStudent: classStudent,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	router.GET("/hello", h.Hello())
	router.POST("/class/create", middlewares.AdminOnly(), h.CreateClass())
	router.GET("/class/search", h.SearchClassByName())
	router.POST("/class/:classId/students", middlewares.AdminOnly(), h.EnrollStudent())
	router.DELETE("/class/:classID/students/:studentID", middlewares.AdminOnly(), h.DeleteStudentFromClass())

	router.POST("/student/create", middlewares.AdminOnly(), h.CreateStudent())
}

func (h *Handler) ConfigRouteAuth(router *gin.RouterGroup) {
	router.POST("/login", h.Login())
	router.POST("/register", h.Register())
}
