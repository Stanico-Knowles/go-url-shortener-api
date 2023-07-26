package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	INTERNAL_SERVER_ERROR string = "An unexpected error occurred"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": INTERNAL_SERVER_ERROR,
				})
			}
		}()
		ctx.Next()
	}
}
