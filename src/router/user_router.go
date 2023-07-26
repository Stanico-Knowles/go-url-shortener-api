package router

import (
	"go-url-shortener-api/src/middlewares"
	"go-url-shortener-api/src/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRouter(api *gin.RouterGroup, db *gorm.DB) {
	repo := user.NewUserRepo(db)
	service := user.NewUserService(repo)
	api.POST("/user/new", user.NewUserController(service).CreateUser)
	api.GET("/user/current-user", middlewares.AuthMiddleware(), user.NewUserController(service).GetCurrentUser)
}
