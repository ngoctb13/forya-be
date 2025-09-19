package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	classStudentUC "github.com/ngoctb13/forya-be/internal/domains/class_student/usecases"
	courseUC "github.com/ngoctb13/forya-be/internal/domains/course/usecases"
	studentUC "github.com/ngoctb13/forya-be/internal/domains/student/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Handler struct {
	user         *userUC.User
	class        *classUC.Class
	student      *studentUC.Student
	classStudent *classStudentUC.ClassStudent
	course       *courseUC.Course
}

func NewHandler(user *userUC.User, class *classUC.Class, student *studentUC.Student, classStudent *classStudentUC.ClassStudent, course *courseUC.Course) *Handler {
	return &Handler{
		user:         user,
		class:        class,
		student:      student,
		classStudent: classStudent,
		course:       course,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	// class
	router.POST("/class/create", middlewares.AdminOnly(), h.CreateClass())
	router.GET("/class/search", h.SearchClassByName())
	router.POST("/class/:classId/students", middlewares.AdminOnly(), h.EnrollStudent())
	router.DELETE("/class/:classId/students/:studentId", middlewares.AdminOnly(), h.DeleteStudentFromClass())

	// student
	router.POST("/student/create", middlewares.AdminOnly(), h.CreateStudent())

	// course
	router.POST("/course/create", middlewares.AdminOnly(), h.CreateCourse())
}

func (h *Handler) ConfigRouteAuth(router *gin.RouterGroup) {
	router.POST("/login", h.Login())
	router.POST("/register", h.Register())
}
