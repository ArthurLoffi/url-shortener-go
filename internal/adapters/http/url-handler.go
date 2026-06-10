package http

import (
	"net/http"
	"strconv"
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

func (h *UrlHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint64, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idUint := uint(idUint64) 
	url, err := h.service.GetByID(c.Request.Context(), idUint)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Url": url,
	})
}

func (h *UrlHandler) GetByShortCode(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Code should be not empty",
		})
	}

	url, err := h.service.GetByShortCode(c.Request.Context(), code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Url": url,
	})
}

func (h *UrlHandler) GetByUserID(c *gin.Context) {
	idUser := c.Param("id")
	idUserUint64, err := strconv.ParseUint(idUser, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	idUserUint := uint(idUserUint64)
	url, err := h.service.GetByUserID(c.Request.Context(), idUserUint)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Url": url,
	})
}