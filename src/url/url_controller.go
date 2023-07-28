package urlshortener

import (
	"go-url-shortener-api/src/middlewares"
	shortenerattributes "go-url-shortener-api/src/url/attributes"
	urlenums "go-url-shortener-api/src/url/enums"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type urlShortenerController struct {
	service URLShortenerService
}

type URLShortenerController interface {
	CreateShortURL(ctx *gin.Context)
	GetOriginalURL(ctx *gin.Context)
	GetURLSByUserID(ctx *gin.Context)
}

func NewURLShortenerController(service URLShortenerService) URLShortenerController {
	return &urlShortenerController{
		service: service,
	}
}

func (controller *urlShortenerController) CreateShortURL(ctx *gin.Context) {
	var url shortenerattributes.CreateShortURLAttributes
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": urlenums.INVALID_REQUEST})
		return
	}
	if validationErrors := controller.service.ValidateInputURL(url.LongURL); validationErrors != (middlewares.ErrorResponse{}) {
		ctx.JSON(validationErrors.Status, gin.H{"error": validationErrors.Message})
		return
	}
	if userId := ctx.GetString("userId"); userId != "" {
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
	url, err := controller.service.GetOriginalURL(ctx.Param("shortUrl"))
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, url.OriginalUrl)
}

func (controller *urlShortenerController) GetURLSByUserID(ctx *gin.Context) {
	var pageSize int
	var pageNumber int
	if pageSize, _ = strconv.Atoi(ctx.Query("pageSize")); pageSize > 100 || pageSize < 1 {
		pageSize = 10
	}
	if pageNumber, _ = strconv.Atoi(ctx.Query("pageNumber")); pageNumber < 1 {
		pageNumber = 1
	}
	urls, err := controller.service.GetURLSByUserID(ctx.GetString("userId"), pageSize, pageNumber)
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"urls": urls})
}
