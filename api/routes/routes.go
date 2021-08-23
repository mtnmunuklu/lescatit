package routes

import (
	"Lescatit/api/middlewares"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route provides the route instance for routing operation.
type Route struct {
	Method       string
	Path         string
	Handler      http.HandlerFunc
	AuthRequired bool
}

// Install registers a new route with a matcher for the URL path.
func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		if route.AuthRequired {
			router.
				HandleFunc(route.Path, middlewares.LogRequests(
					middlewares.Authenticate(route.Handler),
				)).
				Methods(route.Method)
		} else {
			router.
				HandleFunc(route.Path, middlewares.LogRequests(route.Handler)).
				Methods(route.Method)
		}
	}
}

// WithCORS provides Cross-Origin Resource Sharing middleware.
func WithCORS(router *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "Content-Type", "Accept", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	return handlers.CORS(headers, origins, methods)(router)
}
