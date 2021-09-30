package util

import (
	"Lescatit/pb"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/url"
	"strconv"
)

// Contains error codes for categorizer service.
var (
	ErrEmptyData                  = errors.New("data can't be empty")
	ErrEmptyURLs                  = errors.New("urls can't be empty")
	ErrInvalidURL                 = errors.New("invalid url")
	ErrInvalidCategorizationModel = errors.New("invalid categorization model")
	ErrInvalidCount               = errors.New("invalid count")
	ErrFailedModelCreate          = errors.New("failed to create categorization model")
	ErrFailedModelLearn           = errors.New("failed to learn categorization model")
	ErrFailedModelSave            = errors.New("failed to save categorization model to database")
	ErrFailedModelGet             = errors.New("failed to get categorization model from database")
	ErrFailedModelRead            = errors.New("failed to read categorization model")
	ErrFailedModelUpdate          = errors.New("failed to update categorization model")
	ErrFailedModelDelete          = errors.New("failed to delete categorization model")
	ErrFailedModelFind            = errors.New("failed to find model")
	ErrEmptyModelName             = errors.New("model name can't be empty")
	ErrEmptyModelNames            = errors.New("model names can't be empty")
	ErrEmptyModelCategory         = errors.New("model category can't be empty")
	ErrEmptyModelCategories       = errors.New("model categories can't be empty")
)

// ValidateURLs validates if it's a real url.
func ValidateURL(reqURL string) error {
	_, err := url.ParseRequestURI(reqURL)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}

// ValidateURLs validates if it's a real url.
func ValidateURLs(urls []*pb.CategorizeURLRequest) error {
	if len(urls) == 0 {
		return ErrEmptyURLs
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

// ValidateCategories validates the category count.
func ValidateCategories(categories []string) error {
	if len(categories) == 0 {
		return ErrEmptyModelCategories
	}
	return nil
}

// ValidateNames validates the names count.
func ValidateNames(names []string) error {
	if len(names) == 0 {
		return ErrEmptyModelNames
	}
	return nil
}

// ValidateNames validates the names count.
func ValidateData(data string) error {
	if len(data) == 0 {
		return ErrEmptyData
	}
	return nil
}

// GenerateRandomFileName generates a random filename
func GenerateRandomFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return prefix + hex.EncodeToString(randBytes) + suffix
}
