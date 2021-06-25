package validators

import (
	"CWS/pb"
	"errors"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidUserId     = errors.New("invalid userId")
	ErrEmptyName         = errors.New("name can't be empty")
	ErrEmptyEmail        = errors.New("email can't be empty")
	ErrEmptyPassword     = errors.New("password can't be empty")
	ErrEmailAlreadyExist = errors.New("email already exist")
)

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

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
