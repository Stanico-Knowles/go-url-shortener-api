package middlewares

import (
	"fmt"
	"go-url-shortener-api/src/redis"

	"github.com/gin-gonic/gin"
)

// Gets the user info from the session if the user is logged in.
func GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie("sid")
		if err != nil {
			ctx.Next()
			return
		}
		session, err := redis.GetItem(fmt.Sprintf("session:%s", sessionId))
		if err != nil {
			ctx.Next()
			return
		}
		if session == "" {
			ctx.Next()
			return
		}
		ctx.Set("userId", session)
		ctx.Next()
	}
}
