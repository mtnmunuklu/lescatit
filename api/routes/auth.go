package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewAuthRoutes provides the routing process for authentication.
func NewAuthRoutes(authHandlers handlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Path:    "/api/signup",
			Method:  http.MethodPut,
			Handler: authHandlers.SignUp,
		},
		{
			Path:    "/api/signin",
			Method:  http.MethodPost,
			Handler: authHandlers.SignIn,
		},
		{
			Path:         "/api/user",
			Method:       http.MethodGet,
			Handler:      authHandlers.GetUser,
			AuthRequired: true,
		},
		{
			Path:         "/api/user",
			Method:       http.MethodDelete,
			Handler:      authHandlers.DeleteUser,
			AuthRequired: true,
		},
		{
			Path:         "/api/user_rc",
			Method:       http.MethodPost,
			Handler:      authHandlers.ChangeUserRole,
			AuthRequired: true,
		},
		{
			Path:         "/api/user_pu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateUserPassword,
			AuthRequired: true,
		},
		{
			Path:         "/api/user_eu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateUserEmail,
			AuthRequired: true,
		},
		{
			Path:         "/api/user_nu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateUserName,
			AuthRequired: true,
		},
		{
			Path:         "/api/users",
			Method:       http.MethodGet,
			Handler:      authHandlers.GetUsers,
			AuthRequired: true,
		},
	}
}
