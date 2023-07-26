package user

import (
	url "go-url-shortener-api/src/url"
	"time"

	"github.com/google/uuid"
)

// User is a struct that represents a user in the database
type User struct {
	ID        uuid.UUID                   `gorm:"type:char(36);primaryKey"`
	Email     string                      `gorm:"unique;not null"`
	Password  string                      `gorm:"size:255;not null"`
	FirstName string                      `gorm:"size:50;not null"`
	LastName  string                      `gorm:"size:50;not null"`
	URLS      []url.URL `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
