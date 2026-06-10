package http

import (
	"net/http"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	service *services.UrlService
}

func NewUrlHandler(service *services.UrlService) *UrlHandler {
	return &UrlHandler{
		service: service,
	}
}

func (h *UrlHandler) Create(c *gin.Context) {
	var url domain.Url

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Create(c.Request.Context(), &url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": url,
	})
}