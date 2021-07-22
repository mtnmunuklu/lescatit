package security

import (
	"encoding/base64"
	"errors"
)

var (
	ErrBase64Decode = errors.New("invalid base64 data")
)

func Base64Encode(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(url))
}

func Base64Decode(base64Url string) (string, error) {
	url, err := base64.StdEncoding.DecodeString(base64Url)
	if err != nil {
		return "", ErrBase64Decode
	}
	return string(url), nil
}
