package urlshortener

import (
	"errors"
	"go-url-shortener-api/src/middlewares"
	shortenerattributes "go-url-shortener-api/src/url_shortener/attributes"
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
		ctx.Error(&gin.Error{
			Type: gin.ErrorTypePublic,
			Err:  errors.New(err.Error()),
		})
		return
	}
	validationErrors := controller.service.ValidateInputURL(url.LongURL)
	if validationErrors != (middlewares.ErrorResponse{}) {
		ctx.JSON(validationErrors.Status, gin.H{"error": validationErrors.Message})
		return
	}
	var alias string
	if url.Alias == nil || *url.Alias == "" {
		alias = base64encryptionservice.EncodeBase64(url.LongURL)
	} else {
		alias = *url.Alias
	}
	createdUrl, err := controller.service.CreateShortURL(&url, alias)
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
