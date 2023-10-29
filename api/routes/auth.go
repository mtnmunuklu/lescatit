package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/handlers"
)

// NewAuthRoutes provides the routing process for authentication.
func NewAuthRoutes(authHandlers handlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Method: http.MethodPut,
			Path:   "/signup",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.SignUp(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/signin",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.SignIn(c)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/user",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.GetUser(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodDelete,
			Path:   "/user",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.DeleteUser(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/user_rc",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.ChangeUserRole(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/user_pu",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.UpdateUserPassword(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/user_eu",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.UpdateUserEmail(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/user_nu",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.UpdateUserName(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/users",
			Handler: func(c *fiber.Ctx) error {
				return authHandlers.GetUsers(c)
			},
			AuthRequired: true,
		},
	}
}
