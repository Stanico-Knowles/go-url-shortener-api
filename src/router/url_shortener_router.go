package router

import (
	"go-url-shortener-api/src/middlewares"
	urlshortener "go-url-shortener-api/src/url"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitURLShortenerRouter(api *gin.RouterGroup, db *gorm.DB) {
	repo := urlshortener.NewURLShortenerRepo(db)
	service := urlshortener.NewURLShortenerService(repo)
	api.POST("/shorten", middlewares.GetUserInfo(), urlshortener.NewURLShortenerController(service).CreateShortURL)
	api.GET("/:shortUrl", urlshortener.NewURLShortenerController(service).GetOriginalURL)
	api.GET("/current-user/urls", middlewares.AuthMiddleware(), urlshortener.NewURLShortenerController(service).GetURLSByUserID)
}
