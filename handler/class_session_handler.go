package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

func (h *Handler) CreateClassSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.CreateClassSessionRequest{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.ClassID == "" || req.HeldAt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "class_id or held_at is empty"})
			return
		}

		ha, err := time.Parse(time.RFC3339, req.HeldAt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.classSession.CreateClassSession(c, &inputs.CreateClassSessionInput{
			ClassID: req.ClassID,
			HeldAt:  ha,
		})
		if err != nil {
			log.Printf("CreateClassSessionUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "create class session successfully"})
	}
}
