package util

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/security"
)

// Contains error codes for api.
var (
	ErrEmptyBody    = errors.New("body can't be empty")
	ErrEmptyHeader  = errors.New("header can't be empty")
	ErrExistURL     = errors.New("url already exist")
	ErrUnauthorized = errors.New("unauthorized operation")
)

// Error
type JError struct {
	Error string `json:"error"`
}

// WriteAsJSON writes the response in JSON format.
func WriteAsJSON(c *fiber.Ctx, statusCode int, data interface{}) error {
	c.Set("Content-Type", "application/json")
	return c.Status(statusCode).JSON(data)
}

// WriteError writes the error response in JSON format.
func WriteError(c *fiber.Ctx, statusCode int, err error) error {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	return WriteAsJSON(c, statusCode, JError{Error: e})
}

// GetUserIDFromToken returns the user ID from the token.
func GetUserIDFromToken(c *fiber.Ctx) (string, error) {
	token, err := security.ExtractToken(c)
	if err != nil {
		return "", err
	}

	userID, err := security.ValidateToken(token)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// CheckUserIsAdmin checks if user is admin.
func CheckUserIsAdmin(role string) bool {
	return role == "admin"
}
