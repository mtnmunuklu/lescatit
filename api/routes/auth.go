package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewAuthRoutes provides the routing process for authentication.
func NewAuthRoutes(authHandlers handlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path:    "/signup",
			Method:  http.MethodPut,
			Handler: authHandlers.SignUp,
		},
		{
			Path:    "/signin",
			Method:  http.MethodPost,
			Handler: authHandlers.SignIn,
		},
		{
			Path:         "/users",
			Method:       http.MethodGet,
			Handler:      authHandlers.GetUsers,
			AuthRequired: true,
		},
		{
			Path:         "/user",
			Method:       http.MethodGet,
			Handler:      authHandlers.GetUser,
			AuthRequired: true,
		},
		{
			Path:         "/user",
			Method:       http.MethodDelete,
			Handler:      authHandlers.DeleteUser,
			AuthRequired: true,
		},
		{
			Path:         "/user",
			Method:       http.MethodPatch,
			Handler:      authHandlers.ChangeUserRole,
			AuthRequired: true,
		},
		{
			Path:         "/user_nu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateName,
			AuthRequired: true,
		},
		{
			Path:         "/user_pu",
			Method:       http.MethodPatch,
			Handler:      authHandlers.UpdatePassword,
			AuthRequired: true,
		},
		{
			Path:         "/user_eu",
			Method:       http.MethodPatch,
			Handler:      authHandlers.UpdateEmail,
			AuthRequired: true,
		},
	}
}
