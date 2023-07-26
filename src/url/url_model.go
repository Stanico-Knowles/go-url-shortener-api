package urlshortener

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID          uuid.UUID  `gorm:"type:char(36);primaryKey"`
	Alias       string     `gorm:"uniqueIndex;size:100;not null;"`
	OriginalURL string     `gorm:"index;size:255;not null;"`
	UserID      *uuid.UUID `gorm:"type:char(36);"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
