package util

import (
	"encoding/json"
	"errors"
	"net/http"

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

// WriteAsJson provides return the response in json format.
func WriteAsJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

// WriteError provides return the related error in json format.
func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	WriteAsJson(w, statusCode, JError{e})
}

// GetUserIdFromToken provides return the user id in the token.
func GetUserIdFromToken(r *http.Request) (string, error) {
	token, err := security.ExtractToken(r)
	if err != nil {
		return "", err
	}

	userId, err := security.ValidateToken(token)
	if err != nil {
		return "", err
	}

	return userId, nil
}

// CheckUserIsAdmin checks if user is admin.
func CheckUserIsAdmin(role string) bool {
	return role == "admin"
}
