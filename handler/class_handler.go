package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models"
	dm "github.com/ngoctb13/forya-be/internal/domain/models"
)

func (h *Handler) CreateClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.CreateClassRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.class.CreateClass(c, &dm.CreateClassInput{
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

func (h *Handler) SearchClassByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing query param: name"})
			return
		}

		classes, err := h.class.SearchClassByName(c, name)
		if err != nil {
			log.Printf("SearchClassByName fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": classes,
		})
	}
}
