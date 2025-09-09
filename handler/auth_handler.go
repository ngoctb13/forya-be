package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := "login"
		ctx.JSON(http.StatusOK, res)
	}
}
