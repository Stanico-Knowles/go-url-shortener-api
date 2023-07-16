package router

import (
	urlshortener "go-url-shortener-api/src/url_shortener"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitURLShortenerRouter(api *gin.RouterGroup, db *gorm.DB) {
	repo := urlshortener.NewURLShortenerRepo(db)
	service := urlshortener.NewURLShortenerService(repo)
	api.POST("/shorten", urlshortener.NewURLShortenerController(service).CreateShortURL)
	api.GET("/:shortUrl", urlshortener.NewURLShortenerController(service).GetOriginalURL)
}
