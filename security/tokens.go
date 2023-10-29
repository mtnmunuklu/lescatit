package security

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	jwtSecretKey    = []byte(os.Getenv("JWT_SECRET_KEY"))
)

// NewToken creates a new JWT token.
func NewToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
		"iat": time.Now().Unix(),
	})

	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ExtractToken extracts the JWT token from the request header.
func ExtractToken(c *fiber.Ctx) (string, error) {
	// Authorization: Bearer token...
	header := c.Get("Authorization")
	tokenString := strings.TrimPrefix(header, "Bearer ")
	if tokenString == header {
		return "", ErrInvalidToken
	}
	return tokenString, nil
}

// ParseToken parses and verifies the JWT token.
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ValidateToken validates the JWT token and returns the user ID.
func ValidateToken(tokenString string) (string, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["sub"].(string)
		if !ok {
			return "", ErrInvalidToken
		}
		return userId, nil
	}

	return "", ErrInvalidToken
}
