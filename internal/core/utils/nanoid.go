// Package dedicted for nano id generate
// Return a slug for short code in url
package utils

import (
    gonanoid "github.com/jaevor/go-nanoid"
)

func GenerateSlug() (string, error) {
	slug, err := gonanoid.Standard(8)
	if err != nil {
		return "", err
	}
	return slug(), nil
}