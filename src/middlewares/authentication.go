package middlewares

import (
	"fmt"
	"go-url-shortener-api/src/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	UNAUTHORIZED string = "You are not logged in"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie("sid")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": UNAUTHORIZED})
			return
		}
		session, err := redis.GetItem(fmt.Sprintf("session:%s", sessionId))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": UNAUTHORIZED})
			return
		}
		if session == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": UNAUTHORIZED})
			return
		}
		ctx.Set("userId", session)
		ctx.Next()
	}
}
