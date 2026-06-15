package http

import (
	"net/http"
	"strconv"
	"time"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

type ClickHandler struct {
	service *services.ClickService
	urlService *services.UrlService
}

func NewClickHandler(service *services.ClickService, urlService *services.UrlService) *ClickHandler {
	return &ClickHandler{
		service: service,
		urlService: urlService,
	}
}

func (h *ClickHandler) Create(c *gin.Context) {
    urlID := c.Param("urlId")
    urlID64, err := strconv.ParseUint(urlID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "urlId inválido"})
        return
    }

    click := &domain.Click{
        Urlid: uint(urlID64),
        IPAddress: c.ClientIP(),
        ClickedAt: time.Now(),
    }

    if err := h.service.Create(c.Request.Context(), click); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"success": "Click registered"})
}

func (h *ClickHandler) GetByURLID(c *gin.Context) {
	code := c.Param("urlId")

    url, err := h.urlService.GetByShortCode(c.Request.Context(), code)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL não encontrada"})
        return
    }

    clicks, err := h.service.GetByURLID(c.Request.Context(), url.Id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"clicks": clicks})
}

func (h *ClickHandler) CountByURLID(c *gin.Context) {
	code := c.Param("urlId")

    url, err := h.urlService.GetByShortCode(c.Request.Context(), code)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL não encontrada"})
        return
    }

    count, err := h.service.CountByURLID(c.Request.Context(), url.Id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"clicks": count})
}