package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/handler/models/response"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

func (h *Handler) CreateClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.CreateClassRequest{}
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

		err = h.class.CreateClass(c, &inputs.CreateClassInput{
			Name:        req.Name,
			Description: req.Description,
		})

		if err != nil {
			log.Printf("CreateClassUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "create class successfully"})
	}
}

func (h *Handler) ListClassByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.ListClassRequest{}
		if err := c.ShouldBindQuery(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
			return
		}

		classes, pagination, err := h.class.ListClassByName(c, &inputs.SearchClassByNameInput{
			Name:  req.Name,
			Page:  req.Page,
			Limit: req.Limit,
		})

		if err != nil {
			log.Printf("ListClassByNameUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ToSearchClassResponse(classes, pagination))
	}
}

func (h *Handler) EnrollClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.EnrollClassRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if pathClassID := c.Param("classId"); pathClassID != "" {
			req.ClassID = pathClassID
		}

		err := req.Validate()
		if err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err = h.classStudent.EnrollClass(c, &inputs.EnrollClassInput{
			ClassID:    req.ClassID,
			StudentIDs: req.StudentIDs,
		})

		if err != nil {
			log.Printf("EnrollClassUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "enroll class successfully"})
	}
}

func (h *Handler) DeleteStudentFromClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		classID := c.Param("classId")
		studentID := c.Param("studentId")

		if classID == "" || studentID == "" {
			log.Printf("classId and studentId are required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "classId and studentId are required"})
			return
		}

		if err := h.classStudent.DeleteStudentFromClass(c, classID, studentID); err != nil {
			log.Printf("DeleteStudentFromClassUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete student from class successfully"})
	}
}
