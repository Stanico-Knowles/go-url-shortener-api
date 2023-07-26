package authattributes

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}