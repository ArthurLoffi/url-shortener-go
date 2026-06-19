package http

import (
	"net/http"
	"url-shortener-go/internal/core/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// CreateUser godoc
//
// @Summary Create a new user
// @Description Create a user with the provided name
// @Tags users
// @Accept json
// @Produce json
// @Param body body CreateUserRequest true "User data"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var request CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.CreateUser(c.Request.Context(), request.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Success": "User created successfully",
	})
}

// GetUserByName godoc
//
// @Summary Get user by name
// @Description Retrieve a user by name
// @Tags users
// @Produce json
// @Param name path string true "User name"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/users/{name} [get]
func (h *UserHandler) GetUserByName(c *gin.Context) {
	name :=  c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name should be not empty",
		})
		return
	}
	
	user, err := h.service.GetUserByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}