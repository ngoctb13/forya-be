package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/handler/models/response"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

func (h *Handler) CreateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.CreateCourseRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := req.Validate()
		if err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.course.CreateCourse(c, &inputs.CreateCourseInput{
			Name:            req.Name,
			Description:     req.Description,
			SessionCount:    req.SessionCount,
			PricePerSession: req.PricePerSession,
		})

		if err != nil {
			log.Printf("CreateCourseUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "create course successfully"})
	}
}

func (h *Handler) EnrollCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.EnrollCourseRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pathCourseId := c.Param("courseId"); pathCourseId != "" {
			req.CourseID = pathCourseId
		}

		err := req.Validate()
		if err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.courseStudent.CreateCourseStudents(c, &inputs.CreateCourseStudentsInput{
			CourseID:   req.CourseID,
			StudentIDs: req.StudentIDs,
		})

		if err != nil {
			log.Printf("CreateCourseStudentsUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "enroll course successfully"})
	}
}

func (h *Handler) UpdateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		courseId := c.Param("courseId")
		if courseId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid classId"})
			return
		}

		req := &request.UpdateCourseRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fields := inputs.UpdateCourseFields{
			Name:            req.Name,
			Description:     req.Description,
			SessionCount:    req.SessionCount,
			PricePerSession: req.PricePerSession,
		}

		course, err := h.course.UpdateCourse(c, &inputs.UpdateCourseInput{
			CourseID: courseId,
			Fields:   fields,
		})

		if err != nil {
			log.Printf("UpdateCourseUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, course)
	}
}

func (h *Handler) ListCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.ListCoursesRequest{}
		if err := c.ShouldBindQuery(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input, err := req.ValidateAndMap()
		if err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		courses, pagination, err := h.course.ListCourses(c, input)

		if err != nil {
			log.Printf("ListCoursesUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ToListCoursesResponse(courses, pagination))
	}
}
