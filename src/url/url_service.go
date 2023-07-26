package urlshortener

import (
	"go-url-shortener-api/src/middlewares"
	shortenerattributes "go-url-shortener-api/src/url/attributes"
	urlenums "go-url-shortener-api/src/url/enums"
	"net/http"
)

type urlShortenerService struct {
	repo URLShortenerRepo
}

type URLShortenerService interface {
	CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	GetOriginalURL(shortUrl string) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	ValidateInputURL(url string) middlewares.ErrorResponse
}

func NewURLShortenerService(repo URLShortenerRepo) URLShortenerService {
	return &urlShortenerService{
		repo: repo,
	}
}

func (service *urlShortenerService) CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse) {
	existingAlias, _ := service.repo.GetOriginalURL(originalUrl.Alias)
	if existingAlias != nil {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusConflict,
			Message: urlenums.ALIAS_ALREADY_EXISTS,
		}
	}
	existingUrl, _ := service.repo.GetURLSByOriginalURL(originalUrl.LongURL)
	if existingUrl != nil {
		return existingUrl, middlewares.ErrorResponse{}
	}
	newUrl, err := service.repo.CreateShortURL(originalUrl)
	if err != nil {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return newUrl, middlewares.ErrorResponse{}
}

func (service *urlShortenerService) GetOriginalURL(shortUrl string) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse) {
	url, _ := service.repo.GetOriginalURL(shortUrl)
	if url == nil {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: urlenums.URL_NOT_FOUND,
		}
	}
	return url, middlewares.ErrorResponse{}
}

func (service *urlShortenerService) ValidateInputURL(url string) middlewares.ErrorResponse {
	if url == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: urlenums.URL_IS_REQUIRED,
		}
	}
	return middlewares.ErrorResponse{}
}
