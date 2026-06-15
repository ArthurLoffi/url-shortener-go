package main

import (
	"url-shortener-go/docs"
	"url-shortener-go/internal/adapters/database"

	"github.com/gin-gonic/gin"
)

// @title           URL Shortener
// @version         1.0
// @description     API focused on being a URL shortener.
// @host            localhost:8080
// @BasePath        /api
// @in header
// @name Authorization
func main() {
	r := gin.New()

	database.Connect()

	r.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/api"

	SetupRoutes(r)
	r.Run(":8080")
}
