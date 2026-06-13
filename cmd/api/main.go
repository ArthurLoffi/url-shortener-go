package main

import (
	"url-shortener-go/docs"
	"url-shortener-go/internal/adapters/database"
	https "url-shortener-go/internal/adapters/http"
	"url-shortener-go/internal/adapters/repository"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	urlRepo := repository.NewUrlRepository(database.Database)
	urlService := services.NewUrlService(urlRepo)
	urlHandler := https.NewUrlHandler(urlService)

	userRepo := repository.NewUserRepository(database.Database)
	userService := services.NewUserService(userRepo)
	userHandler := https.NewUserHandler(userService)

	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/api/healthy", healthyHandler)
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	url := r.Group("/api/urls")
	{
		url.POST("/", urlHandler.CreateUrl)
		url.GET("/:id", urlHandler.GetByID)
		url.GET("/short/:code", urlHandler.GetByShortCode)
		url.GET("/user/:id", urlHandler.GetByUserID)
	}
	r.GET("/:code", urlHandler.Redirect)

	user := r.Group("/api/users")
	{
		user.POST("/", userHandler.CreateUser)
		user.GET("/:name", userHandler.GetUserByName)
	}

	r.Run(":8080")
}

func healthyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}
