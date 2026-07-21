package main

import (
	"log"
	"url-shortener-go/internal/adapters/outbound/cache"
	"url-shortener-go/internal/adapters/outbound/database"
	https "url-shortener-go/internal/adapters/inbound/http"
	"url-shortener-go/internal/adapters/outbound/repository"
	"url-shortener-go/internal/core/services"
	"url-shortener-go/internal/middleware"

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

	r.GET("/healthy", healthyHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	protect := r.Group("/api")
	protect.Use(middleware.Auth())

	url := protect.Group("/urls")
	{
		url.POST("/", urlHandler.CreateUrl)
		url.GET("/:id", urlHandler.GetByID)
		url.GET("/short/:code", urlHandler.GetByShortCode)
		url.GET("/user/:id", urlHandler.GetByUserID)
	}
	r.GET("/:code", urlHandler.Redirect)

	user := protect.Group("/users")
	{
		user.POST("/", userHandler.CreateUser)
		user.GET("/:name", userHandler.GetUserByName)
	}
	r.POST("/login", userHandler.Login)

	clicks := protect.Group("/clicks")
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