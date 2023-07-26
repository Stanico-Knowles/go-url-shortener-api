package shortenerattributes

type CreateShortURLAttributes struct {
	LongURL string `json:"longUrl"`
	Alias   string `json:"alias"`
	UserID  string `json:"userId"`
}
