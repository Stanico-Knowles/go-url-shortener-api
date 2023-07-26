package router

import (
	"go-url-shortener-api/src/auth"
	"go-url-shortener-api/src/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRouter(api *gin.RouterGroup, db *gorm.DB) {
	repo := user.NewUserRepo(db)
	service := auth.NewAuthService(repo)
	api.POST("/auth/login", auth.NewAuthController(service).Login)
	api.GET("/auth/session", auth.NewAuthController(service).Authenticate)
	api.DELETE("/auth/logout", auth.NewAuthController(service).Logout)
}
