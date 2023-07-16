package migrations

import (
	urlshortener "go-url-shortener-api/src/url_shortener"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(urlshortener.URLShortener{})
}
