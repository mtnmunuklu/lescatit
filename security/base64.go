package security

import (
	"encoding/base64"
	"errors"
)

var (
	ErrBase64Decode = errors.New("invalid base64 data")
)

// Base64Encode provides base64 encoding.
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode provides base64 decoding.
func Base64Decode(base64Data string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", ErrBase64Decode
	}

	return string(data), nil
}
