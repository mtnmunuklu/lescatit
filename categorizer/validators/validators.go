package validators

import (
	"errors"
	"net/url"
)

// Contains error codes for categorizer service.
var (
	ErrEmptyData  = errors.New("data can't be empty")
	ErrInvalidURL = errors.New("invalid url")
)

// ValidateURLs validates if it's a real url.
func ValidateURL(reqURL string) error {
	_, err := url.ParseRequestURI(reqURL)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}
