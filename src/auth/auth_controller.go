package auth

import (
	"fmt"
	authattributes "go-url-shortener-api/src/auth/attributes"
	authenums "go-url-shortener-api/src/auth/enums"
	"go-url-shortener-api/src/middlewares"
	"go-url-shortener-api/src/redis"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type authController struct {
	service AuthService
}

type AuthController interface {
	Login(ctx *gin.Context)
	Authenticate(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func NewAuthController(service AuthService) AuthController {
	return &authController{
		service: service,
	}
}

func (controller *authController) Login(ctx *gin.Context) {
	var credentials authattributes.LoginDTO
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": authenums.INVALID_REQUEST})
		return
	}
	validationErrors := controller.service.ValidateLogin(&credentials)
	if validationErrors != (middlewares.ErrorResponse{}) {
		ctx.JSON(validationErrors.Status, gin.H{"error": validationErrors.Message})
		return
	}
	userData, err := controller.service.Login(&credentials)
	if err != (middlewares.ErrorResponse{}) {
		ctx.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	var sessionId uuid.UUID = uuid.New()
	redis.SetItem(fmt.Sprintf("session:%s", sessionId), userData)
	ctx.SetCookie("sid", sessionId.String(), 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"user": userData})
}

func (controller *authController) Authenticate(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("sid")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": middlewares.UNAUTHORIZED})
		return
	}
	session, err := redis.GetItem(fmt.Sprintf("session:%s", sessionId))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": middlewares.UNAUTHORIZED})
		return
	}
	if session == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": middlewares.UNAUTHORIZED})
		return
	}
	ctx.Status(http.StatusOK)
}

func (controller *authController) Logout(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("sid")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": middlewares.UNAUTHORIZED})
		return
	}
	redis.DeleteItem(fmt.Sprintf("session:%s", sessionId))
	ctx.SetCookie("sid", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "You are logged out"})
}
