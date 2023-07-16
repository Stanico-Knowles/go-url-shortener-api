package shortenerattributes

type ShortUrlResponseAttributes struct {
	Alias       string `json:"alias"`
	OriginalUrl string `json:"originalUrl"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
