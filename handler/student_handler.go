package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models"
	dm "github.com/ngoctb13/forya-be/internal/domain/models"
)

func (h *Handler) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.CreateStudentRequest{}
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

		err = h.student.CreateStudent(c, &dm.CreateStudentInput{
			FullName:          req.FullName,
			Age:               req.Age,
			PhoneNumber:       req.PhoneNumber,
			ParentPhoneNumber: req.ParentPhoneNumber,
			Note:              req.Note,
		})

		if err != nil {
			log.Printf("CreateStudentUsescase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "create student successfully"})
	}
}
