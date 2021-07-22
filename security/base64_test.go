package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	base64Url := Base64Encode("http://example.com")
	assert.NotEmpty(t, base64Url)
	assert.Equal(t, base64Url, "aHR0cDovL2V4YW1wbGUuY29t")
	assert.NotEqual(t, base64Url, "aHR0cDovL2V4YW1wbGUuY29m")
}

func TestBase64Decode(t *testing.T) {
	base64Url := Base64Encode("http://example.com")
	assert.NotEmpty(t, base64Url)

	url, err := Base64Decode(base64Url)
	assert.NoError(t, err)
	assert.NotEmpty(t, url)
	assert.Equal(t, url, "http://example.com")
	assert.NotEqual(t, url, "https://example.com")
}
