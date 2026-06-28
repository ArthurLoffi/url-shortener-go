package http

import (
	"net/http"
	"strconv"
	"time"
	"url-shortener-go/internal/core/domain"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

// Struct criada somente para o swagger
type CreateUrlRequest struct {
    OriginalUrl string `json:"original_url" example:"https://google.com"`
}

type UrlHandler struct {
	service *services.UrlService
	clickService *services.ClickService
}

func NewUrlHandler(service *services.UrlService, clickService *services.ClickService) *UrlHandler {
	return &UrlHandler{
		service: service,
		clickService: clickService,
	}
}

// CreateUrl godoc
//
// @Summary Create a new URL
// @Description Create a shortened URL and store it in the database
// @Security BearerAuth
// @Tags urls
// @Accept json
// @Produce json
// @Param body body CreateUrlRequest true "URL data"
// @Success 201 {object} domain.Url
// @Failure 400
// @Failure 500
// @Router /api/urls [post]
func (h *UrlHandler) CreateUrl(c *gin.Context) {
	var url domain.Url

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, exists := c.Get("id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Unauthorized",
		})
	}

	url.UserID = userID.(uint)

	if err := h.service.CreateUrl(c.Request.Context(), &url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": url,
	})
}

// Redirect godoc
//
// @Summary Redirect to original URL
// @Description Redirects the user to the original URL associated with the short code
// @Security BearerAuth
// @Tags urls
// @Param code path string true "Short URL code"
// @Success 301 "Redirect"
// @Failure 404
// @Failure 410
// @Router /{code} [get]
func (h *UrlHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	url, err := h.service.Redirect(c.Request.Context(), code)
    if err != nil {
        if err.Error() == "url expirada" {
            c.JSON(http.StatusGone, gin.H{"error": "URL expirada"}) // 410
            return
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "URL não encontrada"}) // 404
        return
    }

	click := &domain.Click{
        Urlid: url.Id,
        IPAddress: c.ClientIP(),
        ClickedAt: time.Now(),
    }
    h.clickService.Create(c.Request.Context(), click)

	count, err := h.clickService.CountByURLID(c.Request.Context(), url.Id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
        return
    }

	err = h.service.UpdateClickCount(c.Request.Context(), url.Id, uint(count))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
        return
	}

    c.Redirect(http.StatusMovedPermanently, url.OriginalUrl)
}

// GetByID godoc
//
// @Summary Get URL by ID
// @Description Retrieve a URL by its database ID
// @Security BearerAuth
// @Tags urls
// @Produce json
// @Param id path int true "URL ID"
// @Success 200 {object} domain.Url
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/urls/{id} [get]
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

// GetByShortCode godoc
//
// @Summary Get URL by short code
// @Description Retrieve a URL using its generated short code
// @Security BearerAuth
// @Tags urls
// @Produce json
// @Param code path string true "Short URL code"
// @Success 200 {object} domain.Url
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/urls/short/{code} [get]
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

// GetByUserID godoc
//
// @Summary Get URLs by user ID
// @Description Retrieve all URLs associated with a user
// @Security BearerAuth
// @Tags urls
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} domain.Url
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/urls/user/{id} [get]
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