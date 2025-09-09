package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Handler struct {
	user *usecases.User
}

func NewHandler(user *usecases.User) *Handler {
	return &Handler{
		user: user,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	router.GET("/hello", h.Hello())
}

func (h *Handler) ConfigRouteAuth(router *gin.RouterGroup) {
	router.POST("/login", h.Login())
	router.POST("/register", h.Register())
}
