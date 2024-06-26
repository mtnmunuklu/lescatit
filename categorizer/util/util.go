package util

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/url"
	"strconv"

	"github.com/mtnmunuklu/lescatit/pb"
)

// Contains error codes for categorizer service.
var (
	ErrInvalidURL                 = errors.New("invalid url")
	ErrInvalidCategorizationModel = errors.New("invalid categorization model")
	ErrInvalidCount               = errors.New("invalid count")
	ErrEmptyData                  = errors.New("data can't be empty")
	ErrEmptyURLs                  = errors.New("urls can't be empty")
	ErrEmptyModelName             = errors.New("model name can't be empty")
	ErrEmptyModelNames            = errors.New("model names can't be empty")
	ErrEmptyModelCategory         = errors.New("model category can't be empty")
	ErrEmptyModelCategories       = errors.New("model categories can't be empty")
	ErrCreateModel                = errors.New("model could not be created")
	ErrLearnModel                 = errors.New("model could not be learned")
	ErrSaveModel                  = errors.New("model could not be saved")
	ErrGetModel                   = errors.New("model could not be fetched")
	ErrReadModel                  = errors.New("model could not be read")
	ErrUpdateModel                = errors.New("model could not be updated")
	ErrDeleteModel                = errors.New("model could not be deleted")
	ErrSerializeClassifier        = errors.New("classifier could not be serialized")
	ErrDeserializeClassifer       = errors.New("classifier could not be deserialized")
	ErrDecodeData                 = errors.New("base64 data could not be decoded")
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

// / ValidateCategory validates the category.
func ValidateCategory(category string) error {
	if category == "" {
		return ErrEmptyModelCategory
	}

	return nil
}

// ValidateCategories validates the category count.
func ValidateCategories(categories []string) error {
	if len(categories) == 0 {
		return ErrEmptyModelCategories
	}

	return nil
}

// ValidateName validates the names.
func ValidateName(data string) error {
	if data == "" {
		return ErrEmptyModelName
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

// GenerateRandomFileName generates a random filename.
func GenerateRandomFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	return prefix + hex.EncodeToString(randBytes) + suffix
}
