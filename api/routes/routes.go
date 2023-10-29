package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/middlewares"
)

// Route provides the route instance for routing operation.
type Route struct {
	Method       string
	Path         string
	Handler      fiber.Handler
	AuthRequired bool
}

// Install registers a new route with a matcher for the URL path.
func Install(app *fiber.App, routeList []*Route) {
	for _, route := range routeList {
		handler := route.Handler
		if route.AuthRequired {
			handler = middlewares.Authenticate(handler)
		}
		app.Add(route.Method, route.Path, handler)
	}
}
