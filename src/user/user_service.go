package user

import (
	"go-url-shortener-api/src/middlewares"
	userattributes "go-url-shortener-api/src/user/atrributes"
	userenums "go-url-shortener-api/src/user/enums"
	"go-url-shortener-api/src/utils/hash/bcrypt"
	"net/http"
	"net/mail"
)

type userService struct {
	repo UserRepo
}

type UserService interface {
	CreateUser(user *userattributes.CreateUserDTO) (*userattributes.UserDTO, middlewares.ErrorResponse)
	GetUserByID(id string) (*userattributes.UserDTO, middlewares.ErrorResponse)
	ValidateUser(user *userattributes.CreateUserDTO) middlewares.ErrorResponse
}

func NewUserService(repo UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

func (service *userService) CreateUser(user *userattributes.CreateUserDTO) (*userattributes.UserDTO, middlewares.ErrorResponse) {
	existingUserWithEmail, err := service.repo.GetCountOfUsersByEmail(user.Email)
	if err != nil {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: userenums.SOMETHING_WENT_WRONG,
		}
	}
	if existingUserWithEmail > 0 {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.USER_WITH_EMAIL_ALREADY_EXISTS,
		}
	}
	hash, _ := bcrypt.Hash(user.Password)
	user.Password = string(hash)
	newUser, _ := service.repo.CreateUser(user)
	return newUser, middlewares.ErrorResponse{}
}

func (service *userService) GetUserByID(id string) (*userattributes.UserDTO, middlewares.ErrorResponse) {
	user, err := service.repo.GetUserByID(id)
	if err != nil {
		return nil, middlewares.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: userenums.USER_NOT_FOUND,
		}
	}
	return user, middlewares.ErrorResponse{}
}

func (service *userService) ValidateUser(user *userattributes.CreateUserDTO) middlewares.ErrorResponse {
	if user.Email == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.EMAIL_IS_REQUIRED,
		}
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.EMAIL_IS_INVALID,
		}
	}
	if user.Password == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.PASSWORD_IS_REQUIRED,
		}
	}
	if user.FirstName == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.FIRST_NAME_IS_REQUIRED,
		}
	}
	if user.LastName == "" {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.LAST_NAME_IS_REQUIRED,
		}
	}
	if len(user.Password) < 8 || len(user.Password) > 20 {
		return middlewares.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: userenums.PASSWORD_LENGTH_INVALID,
		}
	}
	return middlewares.ErrorResponse{}
}
