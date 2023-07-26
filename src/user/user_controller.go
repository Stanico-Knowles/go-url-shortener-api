package user

import (
	"go-url-shortener-api/src/middlewares"
	userattributes "go-url-shortener-api/src/user/atrributes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service UserService
}

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetCurrentUser(ctx *gin.Context)
}

func NewUserController(service UserService) UserController {
	return &userController{
		service: service,
	}
}

func (controller *userController) CreateUser(ctx *gin.Context) {
	var user userattributes.CreateUserDTO
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	validationErrors := controller.service.ValidateUser(&user)
	if validationErrors != (middlewares.ErrorResponse{}) {
		ctx.JSON(validationErrors.Status, gin.H{"error": validationErrors.Message})
		return
	}

	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.FirstName = strings.ToLower(strings.TrimSpace(user.FirstName))
	user.LastName = strings.ToLower(strings.TrimSpace(user.LastName))

	createdUser, err := controller.service.CreateUser(&user)
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (controller *userController) GetCurrentUser(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	user, err := controller.service.GetUserByID(userId)
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
