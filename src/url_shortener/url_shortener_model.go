package urlshortener

import (
	"time"

	"github.com/google/uuid"
)

type URLShortener struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Alias       string    `gorm:"uniqueIndex;size:100;not null;"`
	OriginalURL string    `gorm:"uniqueIndex;size:255;not null;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
