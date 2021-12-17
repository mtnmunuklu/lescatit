package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHash tests hashing operation.
func TestHash(t *testing.T) {
	hash := Hash("http://example.com")
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, hash, "1sdasd324wda")
	assert.Equal(t, hash, "f0e6a6a97042a4f1f1c87f5f7d44315b2d852c2df5c7991cc66241bf7072d1c4")
}
