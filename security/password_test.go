package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBase64Encode tests password encrypt operation.
func TestEncryptPassword(t *testing.T) {
	pass, err := EncryptPassword("123456789")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
	assert.Len(t, pass, 60)
}

// TestVerifyPassword tests password verify operation.
func TestVerifyPassword(t *testing.T) {
	pass, err := EncryptPassword("123456789")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)

	assert.NoError(t, VerifyPassword(pass, "123456789"))
	assert.Error(t, VerifyPassword(pass, "1234567"))
	assert.Error(t, VerifyPassword(pass, pass))
	assert.Error(t, VerifyPassword("123456789", pass))
}
