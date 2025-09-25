package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models"
	dm "github.com/ngoctb13/forya-be/internal/domain/models"
)

func (h *Handler) CreateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.CreateCourseRequest{}
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

		err = h.course.CreateCourse(c, &dm.CreateCourseInput{
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
		req := &models.EnrollCourseRequest{}
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

		err = h.courseStudent.CreateCourseStudents(c, &dm.CreateCourseStudentsInput{
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

		req := &models.UpdateCourseRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		course, err := h.course.UpdateCourse(c, &dm.UpdateCourseInput{
			CourseID: courseId,
			Fields:   req.Fields,
		})

		if err != nil {
			log.Printf("UpdateCourseUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, course)
	}
}
