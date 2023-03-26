package util

import (
	"errors"
	"strings"

	"github.com/mtnmunuklu/lescatit/pb"
)

// Contains error codes for authentication service.
var (
	ErrInvalidUserId         = errors.New("invalid user id")
	ErrEmptyName             = errors.New("name can't be empty")
	ErrEmptyEmail            = errors.New("email can't be empty")
	ErrEmptyNewEmail         = errors.New("new email can't be empty")
	ErrEmptyPassword         = errors.New("password can't be empty")
	ErrEmptyNewPassword      = errors.New("new password can't be empty")
	ErrEmptyUserRole         = errors.New("user role can't be empty")
	ErrExistEmail            = errors.New("email already exist")
	ErrNotFoundEmail         = errors.New("email did not found")
	ErrNotFoundUserId        = errors.New("user id could not be found")
	ErrFailedSignIn          = errors.New("signin failed")
	ErrMismatchedPassword    = errors.New("password did not match")
	ErrCreateUser            = errors.New("user could not be created")
	ErrDeleteUser            = errors.New("user could not be deleted")
	ErrUpdateUser            = errors.New("user could not be updated")
	ErrEncryptPassword       = errors.New("password could not be encrypted")
	ErrNotPerformedOperation = errors.New("operation could not be performed")
)

// ValidateSingnUp validates the user information for user registration process.
func ValidateSignUp(user *pb.SignUpRequest) error {
	if user.GetEmail() == "" {
		return ErrEmptyEmail
	} else if user.GetName() == "" {
		return ErrEmptyName
	} else if user.GetPassword() == "" {
		return ErrEmptyPassword
	}

	return nil
}

// NormalizeEmail normalizes the user email address.
func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
