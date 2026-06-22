package utils

import (
	"os"
	"time"

	"url-shortener-go/internal/middleware"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, name string) (string, error) {
	claims := middleware.Claims{
		Id: userID,
		User: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}