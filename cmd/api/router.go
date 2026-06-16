package main

import (
	"log"
	"url-shortener-go/internal/adapters/cache"
	"url-shortener-go/internal/adapters/database"
	https "url-shortener-go/internal/adapters/http"
	"url-shortener-go/internal/adapters/repository"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	redisClient, err := database.NewRedisClient()
    if err != nil {
        log.Fatal(err)
    }
	
	urlCache := cache.NewUrlCache(redisClient)

	urlRepo := repository.NewUrlRepository(database.Database)
	urlService := services.NewUrlService(urlRepo, urlCache)

	clickRepo := repository.NewClickRepository(database.Database)
	clickService := services.NewClickService(clickRepo)
	clickHandler := https.NewClickHandler(clickService, urlService)
	
	urlHandler := https.NewUrlHandler(urlService, clickService)

	userRepo := repository.NewUserRepository(database.Database)
	userService := services.NewUserService(userRepo)
	userHandler := https.NewUserHandler(userService)

	
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

	clicks := r.Group("/api/clicks")
    {
        clicks.POST("/:urlId", clickHandler.Create)
        clicks.GET("/:urlId", clickHandler.GetByURLID)
        clicks.GET("/:urlId/count", clickHandler.CountByURLID)
    }
}

func healthyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}