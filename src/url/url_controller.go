package urlshortener

import (
	"go-url-shortener-api/src/middlewares"
	shortenerattributes "go-url-shortener-api/src/url/attributes"
	base64encryptionservice "go-url-shortener-api/src/utils/hash/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlShortenerController struct {
	service URLShortenerService
}

type URLShortenerController interface {
	CreateShortURL(ctx *gin.Context)
	GetOriginalURL(ctx *gin.Context)
}

func NewURLShortenerController(service URLShortenerService) URLShortenerController {
	return &urlShortenerController{
		service: service,
	}
}

func (controller *urlShortenerController) CreateShortURL(ctx *gin.Context) {
	var url shortenerattributes.CreateShortURLAttributes
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	validationErrors := controller.service.ValidateInputURL(url.LongURL)
	if validationErrors != (middlewares.ErrorResponse{}) {
		ctx.JSON(validationErrors.Status, gin.H{"error": validationErrors.Message})
		return
	}
	var alias string = base64encryptionservice.EncodeBase64(url.LongURL)
	if url.Alias == "" {
		url.Alias = alias
	}
	userId := ctx.GetString("userId")
	if userId != "" {
		url.UserID = userId
	}
	createdUrl, err := controller.service.CreateShortURL(&url)
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"urls": createdUrl})
}

func (controller *urlShortenerController) GetOriginalURL(ctx *gin.Context) {
	response, err := controller.service.GetOriginalURL(ctx.Param("shortUrl"))
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"urls": response})
}
