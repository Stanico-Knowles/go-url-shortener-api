package urlshortener

import (
	"go-url-shortener-api/src/middlewares"
	shortenerattributes "go-url-shortener-api/src/url_shortener/attributes"
	urlshortenerenums "go-url-shortener-api/src/url_shortener/enums"
	"log"
	"net/http"
)

type urlShortenerService struct {
	repo URLShortenerRepo
}

type URLShortenerService interface {
	CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes, alias string) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	GetOriginalURL(shortUrl string) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	ValidateInputURL(url string) middlewares.ErrorResponse
}

func NewURLShortenerService(repo URLShortenerRepo) URLShortenerService {
	return &urlShortenerService{
		repo: repo,
	}
}

func (service *urlShortenerService) CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes, alias string) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse) {
	existingUrl, _ := service.repo.GetURLSByOriginalURL(originalUrl.LongURL)

	if existingUrl != nil {
		log.Println(existingUrl)
		return existingUrl, middlewares.ErrorResponse{}
	}
	newUrl, err := service.repo.CreateShortURL(originalUrl, alias)
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
			Message: urlshortenerenums.URL_NOT_FOUND,
		}
	}
	return url, middlewares.ErrorResponse{}
}

func (service *urlShortenerService) ValidateInputURL(url string) middlewares.ErrorResponse {
	if url == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: urlshortenerenums.URL_IS_REQUIRED,
		}
	}
	return middlewares.ErrorResponse{}
}
