package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	classStudentUC "github.com/ngoctb13/forya-be/internal/domains/class_student/usecases"
	courseUC "github.com/ngoctb13/forya-be/internal/domains/course/usecases"
	courseStudentUC "github.com/ngoctb13/forya-be/internal/domains/course_student/usecases"
	studentUC "github.com/ngoctb13/forya-be/internal/domains/student/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Handler struct {
	user          *userUC.User
	class         *classUC.Class
	student       *studentUC.Student
	classStudent  *classStudentUC.ClassStudent
	course        *courseUC.Course
	courseStudent *courseStudentUC.CourseStudent
}

func NewHandler(user *userUC.User,
	class *classUC.Class,
	student *studentUC.Student,
	classStudent *classStudentUC.ClassStudent,
	course *courseUC.Course,
	courseStudent *courseStudentUC.CourseStudent) *Handler {
	return &Handler{
		user:          user,
		class:         class,
		student:       student,
		classStudent:  classStudent,
		course:        course,
		courseStudent: courseStudent,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	// class
	router.POST("/class/create", middlewares.AdminOnly(), h.CreateClass())
	router.GET("/class/search", h.SearchClassByName())
	router.POST("/class/:classId/students", middlewares.AdminOnly(), h.EnrollClass())
	router.DELETE("/class/:classId/student/:studentId", middlewares.AdminOnly(), h.DeleteStudentFromClass())

	// student
	router.POST("/student/create", middlewares.AdminOnly(), h.CreateStudent())
	router.POST("student/import", middlewares.AdminOnly(), h.ImportStudentsCSVFile())

	// course
	router.POST("/course/create", middlewares.AdminOnly(), h.CreateCourse())
	router.POST("/course/:courseId/enroll", middlewares.AdminOnly(), h.EnrollCourse())
}

func (h *Handler) ConfigRouteAuth(router *gin.RouterGroup) {
	router.POST("/login", h.Login())
	router.POST("/register", h.Register())
}
