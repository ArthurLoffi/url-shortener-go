package main

import (
	"log/slog"
	"os"
	"url-shortener-go/docs"
	"url-shortener-go/internal/adapters/database"
	"url-shortener-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// @title           URL Shortener
// @version         1.0
// @description     API focused on being a URL shortener.
// @host            localhost
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.New()

	r.Use(middleware.RateLimitMiddleware(middleware.NewIPRateLimiter()))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    r.Use(gin.Recovery())
    r.Use(middleware.SlogMiddleware(logger))

	database.Connect()
	docs.SwaggerInfo.BasePath = "/"

	SetupRoutes(r)
	r.Run(":8080")
}
