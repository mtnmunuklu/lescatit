package security

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// TestNewToken tests new a token create operation.
func TestNewToken(t *testing.T) {
	id := bson.NewObjectId()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// TestNewTokenPayload tests new a token payload create operation.
func TestNewTokenPayload(t *testing.T) {
	id := bson.NewObjectId()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := NewTokenPayload(token)
	assert.NoError(t, err)
	assert.NotNil(t, payload)
	assert.Equal(t, id.Hex(), payload.UserId)

	tokenExpired := GetTokenExpired(id.Hex())
	payload, err = NewTokenPayload(tokenExpired)
	assert.Error(t, err)
	assert.Nil(t, payload)
}

// GetTokenExpired tests the operation of receiving the token expire duration.
func GetTokenExpired(id string) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5 * -1).Unix(),
		Issuer:    id,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecretKey)
	return tokenString
}
