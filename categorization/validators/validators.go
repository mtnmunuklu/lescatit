package validators

import (
	"errors"
	"net/url"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

// Contains error codes for categorization service.
var (
	ErrInvalidURLId  = errors.New("invalid url id")
	ErrInvalidURL    = errors.New("invalid url")
	ErrEmptyURLs     = errors.New("urls can't be empty")
	ErrEmptyCategory = errors.New("category can't be empty")
	ErrInvalidCount  = errors.New("invalid count")
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

// ValidateCategories validates the category count.
func ValidateCategories(categories []string) error {
	if len(categories) == 0 {
		return ErrEmptyCategory
	}
	return nil
}

// ValidateCount validates if it's a integer count.
func ValidateCount(count string) (int, error) {
	newCount, err := strconv.Atoi(count)
	if err != nil {
		return 0, ErrInvalidCount
	}
	return newCount, nil
}

// ValidateId validates if it's a valid url id.
func ValidateId(id string) error {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidURLId
	}
	return nil
}
