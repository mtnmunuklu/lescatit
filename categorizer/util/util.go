package util

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/url"
)

// Contains error codes for categorizer service.
var (
	ErrEmptyData                  = errors.New("data can't be empty")
	ErrInvalidURL                 = errors.New("invalid url")
	ErrInvalidCategorizationModel = errors.New("invalid categorization model")
	ErrFailedCreateModel          = errors.New("failed to create categorization model")
)

// ValidateURLs validates if it's a real url.
func ValidateURL(reqURL string) error {
	_, err := url.ParseRequestURI(reqURL)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}

// GenerateRandomFileName generates a random filename for use in testing or whatever
func GenerateRandomFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return prefix + hex.EncodeToString(randBytes) + suffix
}
