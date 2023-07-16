package main

import (
	"go-url-shortener-api/src/database"
	"go-url-shortener-api/src/middlewares"
	"go-url-shortener-api/src/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.Connect()
)

func main() {
	defer database.Disconnect(db)

	server := gin.Default()
	server.SetTrustedProxies(nil)
	server.Use(middlewares.CORS())
	server.Use(middlewares.ErrorMiddleware())

	apiV1 := server.Group("/api/v1")
	router.InitURLShortenerRouter(apiV1, db)

	server.Run()
}
