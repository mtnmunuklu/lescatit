package util

import (
	"Lescatit/security"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// Contains error codes for api.
var (
	ErrEmptyBody       = errors.New("body can't be empty")
	ErrEmptyHeader     = errors.New("header can't be empty")
	ErrURLAlreadyExist = errors.New("url already exist")
	ErrUnauthorized    = errors.New("unauthorized")
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

// AuthRequestWithId checks for unauthorized access by comparing the id in the token with the id used in the requested transaction.
func AuthRequestWithId(r *http.Request) (*security.TokenPayload, error) {
	token, err := security.ExtractToken(r)
	if err != nil {
		return nil, err
	}
	payload, err := security.NewTokenPayload(token)
	if err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	if payload.UserId != vars["id"] {
		return nil, ErrUnauthorized
	}
	return payload, nil
}
