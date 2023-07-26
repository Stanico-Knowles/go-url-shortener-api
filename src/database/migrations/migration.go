package migrations

import (
	url "go-url-shortener-api/src/url"
	"go-url-shortener-api/src/user"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(url.URL{})
	db.AutoMigrate(user.User{})
}
