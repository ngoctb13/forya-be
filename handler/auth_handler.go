package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/models"
	dm "github.com/ngoctb13/forya-be/internal/domain/models"
	"github.com/ngoctb13/forya-be/pkg/auth"
	"github.com/ngoctb13/forya-be/utils"
)

func (h *Handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.LoginRequest{}
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
			log.Printf("GetUserByUsername got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		err = utils.ComparePassword(req.Password, user.Password)
		if err != nil {
			log.Printf("CompareHashAndPassword got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := auth.GenerateJWT(user.ID, user.Role)
		if err != nil {
			log.Printf("GenerateJWT got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (h *Handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.CreateUserRequest{}
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

		err = h.user.CreateUser(c, &dm.CreateUserInput{
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
