package auth

import (
	authattributes "go-url-shortener-api/src/auth/attributes"
	authenums "go-url-shortener-api/src/auth/enums"
	"go-url-shortener-api/src/middlewares"
	"go-url-shortener-api/src/user"
	"go-url-shortener-api/src/utils/hash/bcrypt"
	"net/http"
)

type authService struct {
	repo user.UserRepo
}

type AuthService interface {
	Login(credentials *authattributes.LoginDTO) (string, middlewares.ErrorResponse)
	ValidateLogin(credentials *authattributes.LoginDTO) middlewares.ErrorResponse
}

func NewAuthService(repo user.UserRepo) AuthService {
	return &authService{
		repo: repo,
	}
}

func (service *authService) Login(credentials *authattributes.LoginDTO) (string, middlewares.ErrorResponse) {
	user, err := service.repo.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", middlewares.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: authenums.INVALID_CREDENTIALS,
		}
	}
	if !bcrypt.CompareHashAndPlainText(user.Password, credentials.Password) {
		return "", middlewares.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: authenums.INVALID_CREDENTIALS,
		}
	}
	return user.ID, middlewares.ErrorResponse{}
}

func (service *authService) ValidateLogin(credentials *authattributes.LoginDTO) middlewares.ErrorResponse {
	if credentials.Email == "" || credentials.Password == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: authenums.CREDENTIALS_REQUIRED,
		}
	}
	return middlewares.ErrorResponse{}
}
