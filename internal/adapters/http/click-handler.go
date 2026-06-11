package http

import (
	"net/http"
	"strconv"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

type ClickHandler struct {
	service services.ClickService
}

func NewClickHandler(service services.ClickService) *ClickHandler {
	return &ClickHandler{
		service: service,
	}
}

func (h *ClickHandler) Create(c *gin.Context) {
	var request *domain.Click

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Create(c.Request.Context(), request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "User created successfully",
	})
}

func (h *ClickHandler) GetByURLID(c *gin.Context) {
	urlID := c.Param("urlId")
	urlID64, err := strconv.ParseUint(urlID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	urlIDuint := uint(urlID64)

	click, err := h.service.GetByURLID(c.Request.Context(), urlIDuint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"click": click,
	})
}

func (h *ClickHandler) CountByURLID(c *gin.Context) {
	urlID := c.Param("urlId")
	urlID64, err := strconv.ParseUint(urlID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	urlIDuint := uint(urlID64)

	click, err := h.service.CountByURLID(c.Request.Context(), urlIDuint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clicks": click,
	})
}