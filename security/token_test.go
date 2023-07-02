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

// TestValidateToken tests the validation of a token.
func TestValidateToken(t *testing.T) {
	id := bson.NewObjectId()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userId, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, userId)
	assert.Equal(t, id.Hex(), userId)

	tokenExpired := GetTokenExpired(id.Hex())
	userId, err = ValidateToken(tokenExpired)
	assert.Error(t, err)
	assert.Nil(t, userId)
}

// GetTokenExpired tests the operation of receiving the token expire duration.
func GetTokenExpired(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Minute * 5 * -1).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecretKey)
	return tokenString
}
