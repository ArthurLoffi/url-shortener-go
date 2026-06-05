package main

import (
	"url-shortener-go/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
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

	r.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/api/healthy", healthyHandler)
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}


// Healthy godoc
// @Summary      Health check
// @Description  Returns API health status
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]bool
// @Router       /healthy [get]
func healthyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}