package util

import (
	"Lescatit/pb"
	"errors"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// Contains error codes for authentication service.
var (
	ErrInvalidUserId     = errors.New("invalid userId")
	ErrEmptyName         = errors.New("name can't be empty")
	ErrEmptyEmail        = errors.New("email can't be empty")
	ErrEmptyPassword     = errors.New("password can't be empty")
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrSignInFailed      = errors.New("signin failed")
)

// ValidateSingnUp validates the user information for user registration process.
func ValidateSignUp(user *pb.User) error {
	if !bson.IsObjectIdHex(user.Id) {
		return ErrInvalidUserId
	} else if user.Email == "" {
		return ErrEmptyEmail
	} else if user.Name == "" {
		return ErrEmptyName
	} else if user.Password == "" {
		return ErrEmptyPassword
	}
	return nil
}

// NormalizeEmail normalizes the user email address.
func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
