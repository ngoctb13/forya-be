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

		err := h.course.CreateCourse(c, &dm.CreateCourseInput{
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
