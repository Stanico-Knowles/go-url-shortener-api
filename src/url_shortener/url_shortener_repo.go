package urlshortener

import (
	shortenerattributes "go-url-shortener-api/src/url_shortener/attributes"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type urlShortenerRepo struct {
	DB *gorm.DB
}

type URLShortenerRepo interface {
	CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes, alias string) (*shortenerattributes.ShortUrlResponseAttributes, error)
	GetOriginalURL(shortUrl string) (*shortenerattributes.ShortUrlResponseAttributes, error)
	GetURLSByOriginalURL(originalUrl string) (*shortenerattributes.ShortUrlResponseAttributes, error)
}

func NewURLShortenerRepo(db *gorm.DB) URLShortenerRepo {
	return &urlShortenerRepo{
		DB: db,
	}
}

func (repo *urlShortenerRepo) CreateShortURL(originalUrl *shortenerattributes.CreateShortURLAttributes, alias string) (*shortenerattributes.ShortUrlResponseAttributes, error) {
	newUrl := URLShortener{
		ID:          uuid.New(),
		Alias:       alias,
		OriginalURL: originalUrl.LongURL,
	}
	result := repo.DB.Create(&newUrl)
	if result.Error != nil {
		return nil, result.Error
	}
	return toURLResponseDTO(&newUrl), nil
}

func (repo *urlShortenerRepo) GetOriginalURL(shortUrl string) (*shortenerattributes.ShortUrlResponseAttributes, error) {
	var urlShortener URLShortener
	result := repo.DB.Where("alias = ?", shortUrl).First(&urlShortener)
	if result.Error != nil {
		return nil, result.Error
	}
	return toURLResponseDTO(&urlShortener), nil
}

func (repo *urlShortenerRepo) GetURLSByOriginalURL(originalUrl string) (*shortenerattributes.ShortUrlResponseAttributes, error) {
	var urlShortener URLShortener
	result := repo.DB.Where("original_url = ?", originalUrl).First(&urlShortener)
	if result.Error != nil {
		return nil, result.Error
	}
	return toURLResponseDTO(&urlShortener), nil
}

func toURLResponseDTO(urlShortener *URLShortener) *shortenerattributes.ShortUrlResponseAttributes {
	return &shortenerattributes.ShortUrlResponseAttributes{
		Alias:       urlShortener.Alias,
		OriginalUrl: urlShortener.OriginalURL,
		CreatedAt:   urlShortener.CreatedAt.String(),
		UpdatedAt:   urlShortener.UpdatedAt.String(),
	}
}
