package main

import (
	"url-shortener-go/docs"
	"url-shortener-go/internal/adapters/database"

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

    docs.SwaggerInfo.BasePath = "/api"

    r.GET("/api/healthy", healthyHandler)
    r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    r.Run(":8080")
}

func healthyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}