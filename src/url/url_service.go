package urlshortener

import (
	"go-url-shortener-api/src/middlewares"
	shortenerattributes "go-url-shortener-api/src/url/attributes"
	urlenums "go-url-shortener-api/src/url/enums"
	base64encryptionservice "go-url-shortener-api/src/utils/hash/base64"
	"net/http"
)

type urlShortenerService struct {
	repo URLShortenerRepo
}

type URLShortenerService interface {
	CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	GetOriginalURL(shortUrl string) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	GetURLSByUserID(userID string) ([]*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse)
	ValidateInputURL(url string) middlewares.ErrorResponse
}

func NewURLShortenerService(repo URLShortenerRepo) URLShortenerService {
	return &urlShortenerService{
		repo: repo,
	}
}

func (service *urlShortenerService) CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes) (*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse) {
	if originalUrl.Alias != "" {
		if count, _ := service.getCountOfField("alias", originalUrl.Alias); count > 0 {
			return nil, middlewares.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: urlenums.ALIAS_ALREADY_EXISTS,
			}
		}
	} else {
		originalUrl.Alias = base64encryptionservice.EncodeBase64(originalUrl.LongURL)
	}
	if existingUrl, _ := service.repo.GetURLSByOriginalURL(originalUrl.LongURL); existingUrl != nil {
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

func (service *urlShortenerService) GetURLSByUserID(userID string) ([]*shortenerattributes.ShortUrlResponseAttributes, middlewares.ErrorResponse) {
	if userID == "" {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: urlenums.INVALID_REQUEST,
		}
	}
	urls, _ := service.repo.GetURLSByUserID(userID)
	return urls, middlewares.ErrorResponse{}
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

func (service *urlShortenerService) getCountOfField(field string, value string) (int64, error) {
	count, err := service.repo.GetCountOfField(field, value)
	if err != nil {
		return 0, err
	}
	return count, nil
}
