package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/mtnmunuklu/lescatit/api/util"
	"github.com/mtnmunuklu/lescatit/security"
)

// LogRequests provides logging of incoming requests.
func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		log.Printf(`{"proto": "%s", "method": "%s", "route": "%s%s", "request_time": "%v"}`, r.Proto, r.Method, r.Host, r.URL.Path, time.Since(t))
	}
}

// Authenticate provides the authentication process.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := security.ExtractToken(r)
		if err != nil {
			util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
			return
		}
		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("error on parse token:", err.Error())
			util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
			return
		}
		if !token.Valid {
			log.Println("invalid token:", tokenString)
			util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
			return
		}

		next(w, r)
	}
}
