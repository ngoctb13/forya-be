package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models/request"
	"github.com/ngoctb13/forya-be/handler/models/response"
	"github.com/ngoctb13/forya-be/internal/domains/inputs"
	"github.com/ngoctb13/forya-be/utils"
)

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.LoginRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := req.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user, err := h.user.GetUserByUsername(c, req.UserName)
		if err != nil {
			log.Printf("GetUserByUsernameUsecase got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		err = utils.ComparePassword(req.Password, user.Password)
		if err != nil {
			log.Printf("CompareHashAndPassword got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		at, err := h.auth.GenerateAccessToken(user.ID, user.Role)
		if err != nil {
			log.Printf("GenerateAccessTokenUsecase got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		rt, err := h.auth.GenerateRefreshToken(c, user.ID, user.Role)
		if err != nil {
			log.Printf("GenerateRefreshTokenUsecase got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":          response.ToUserResponse(user),
			"token":         at,
			"refresh_token": rt,
		})
	}
}

func (h *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &request.CreateUserRequest{}
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

		hashPwd, err := utils.HashPassword(req.Password)
		if err != nil {
			log.Printf("fail to hash password: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.user.CreateUser(c, &inputs.CreateUserInput{
			Email:    req.Email,
			UserName: req.UserName,
			Password: hashPwd,
		})

		if err != nil {
			log.Printf("CreateUserUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "register successfully"})
	}
}

func (h *Handler) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		at, rt, err := h.auth.RefreshAccessToken(c, req.RefreshToken)
		if err != nil {
			log.Printf("RefreshAccessTokenUsecase got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":         at,
			"refresh_token": rt,
		})
	}
}

func (h *Handler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token is required"})
			return
		}

		err := h.auth.RevokeAccessToken(c, req.RefreshToken)
		if err != nil {
			log.Printf("RevokeAccessTokenUsecase got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
	}
}
