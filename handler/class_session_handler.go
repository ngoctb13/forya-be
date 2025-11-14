package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/handler/models/response"
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
			Name:    req.Name,
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

func (h *Handler) ListClassSessions() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.ListClassSessionsRequest{}
		if err := c.ShouldBindQuery(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input := &inputs.ListClassSessionsInput{
			ClassID: req.ClassID,
			Page:    req.Page,
			Limit:   req.Limit,
		}

		if req.StartTime != nil {
			t, err := time.Parse(time.RFC3339, *req.StartTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			input.StartTime = &t
		}

		if req.EndTime != nil {
			t, err := time.Parse(time.RFC3339, *req.EndTime)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			input.EndTime = &t
		}

		sessions, p, err := h.classSession.ListClassSessions(c, input)
		if err != nil {
			log.Printf("ListClassSessions fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ToListClassSessionsResponse(sessions, p))
	}
}

func (h *Handler) BatchMarkClassSessionAttendance() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.Param("sessionId")
		if sessionID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sessionId"})
			return
		}

		req := &request.BatchMarkClassSessionAttendanceRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert request to input
		attendanceItems := make([]inputs.AttendanceItem, 0, len(req.Attendances))
		for _, att := range req.Attendances {
			attendanceItems = append(attendanceItems, inputs.AttendanceItem{
				CourseStudentID: att.CourseStudentID,
				IsAttended:      att.IsAttended,
			})
		}

		err := h.classSession.BatchMarkAttendance(c, &inputs.BatchMarkClassSessionAttendanceInput{
			ClassSessionID: sessionID,
			Attendances:    attendanceItems,
		})

		if err != nil {
			log.Printf("BatchMarkClassSessionAttendance fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "batch mark attendance successfully",
			"count":   len(attendanceItems),
		})
	}
}
