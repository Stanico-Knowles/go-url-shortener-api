package main

import (
	"fmt"
	"go-url-shortener-api/src/database"
	"go-url-shortener-api/src/middlewares"
	"go-url-shortener-api/src/redis"
	"go-url-shortener-api/src/router"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.Connect()
)

func main() {
	defer database.Disconnect(db)

	redis.ConnectRedis()

	defer redis.CloseRedis()

	server := gin.Default()
	server.SetTrustedProxies(nil)
	server.Use(middlewares.CORS())
	server.Use(middlewares.ErrorMiddleware())

	apiV1 := server.Group("/api/v1")
	router.InitURLShortenerRouter(apiV1, db)
	router.InitAuthRouter(apiV1, db)
	router.InitUserRouter(apiV1, db)

	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
