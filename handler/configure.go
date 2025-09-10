package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/forya-be/handler/middlewares"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
)

type Handler struct {
	user  *userUC.User
	class *classUC.Class
}

func NewHandler(user *userUC.User, class *classUC.Class) *Handler {
	return &Handler{
		user:  user,
		class: class,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	router.GET("/hello", h.Hello())
	router.POST("/create/class", middlewares.AdminOnly(), h.CreateClass())
}

func (h *Handler) ConfigRouteAuth(router *gin.RouterGroup) {
	router.POST("/login", h.Login())
	router.POST("/register", h.Register())
}
