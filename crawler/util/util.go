package util

import (
	"errors"
	"net/url"
)

// Contains error codes for crawler service.
var (
	ErrInvalidURL  = errors.New("invalid url")
	ErrEmptyURLs   = errors.New("urls can't be empty")
	ErrURLNotExist = errors.New("url does not exist")
)

// ValidateURLs validates if it's a real url.
func ValidateURL(reqURL string) error {
	_, err := url.ParseRequestURI(reqURL)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}

// ValidateURLs validates the url count.
func ValidateURLs(urls []string) error {
	if len(urls) == 0 {
		return ErrEmptyURLs
	}
	return nil
}