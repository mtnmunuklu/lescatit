package security

import (
	"crypto/sha256"
	"fmt"
)

// Hash provides hashing.
func Hash(url string) string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	return hash
}
