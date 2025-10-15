package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/handler/models/response"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
)

func (h *Handler) CreateSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.CreateSupplyRequest{}
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

		err = h.supply.CreateSupply(c, &inputs.CreateSupplyInput{
			Name:         req.Name,
			Description:  req.Description,
			Unit:         req.Unit,
			MinThreshold: req.MinThreshold,
		})

		if err != nil {
			log.Printf("CreateSupplyUsescase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "create supply successfully"})
	}
}

func (h *Handler) ListSupplies() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.ListSuppliesRequest{}
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

		out, p, err := h.supply.ListSupplies(c, input)
		if err != nil {
			log.Printf("ListClassSessions fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ToListSuppliesResponse(out, p))
	}
}

func (h *Handler) UpdateSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		supplyId := c.Param("supplyId")
		if supplyId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplyId"})
			return
		}

		req := &request.UpdateStudentRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			log.Printf("validate request error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.supply.UpdateSupply(c, &inputs.UpdateSupplyInput{
			ID:     supplyId,
			Fields: req.Fields,
		})

		if err != nil {
			log.Printf("UpdateSupplyUsecase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "update supply successfully"})
	}
}

func (h *Handler) DeleteSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		supplyID := c.Param("supplyId")
		if supplyID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "supply id is required"})
			return
		}

		err := h.supply.DeleteSupply(c, supplyID)
		if err != nil {
			log.Printf("DeleteSupplyUsescase fail with error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete supply successfully"})
	}
}
