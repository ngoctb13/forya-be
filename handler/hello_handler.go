package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := "hello world"
		ctx.JSON(http.StatusOK, res)
	}
}
