package routes

import (
	"net/http"

	"github.com/mtnmunuklu/lescatit/api/handlers"
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
			Path:         "/user_rc",
			Method:       http.MethodPost,
			Handler:      authHandlers.ChangeUserRole,
			AuthRequired: true,
		},
		{
			Path:         "/user_pu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateUserPassword,
			AuthRequired: true,
		},
		{
			Path:         "/user_eu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateUserEmail,
			AuthRequired: true,
		},
		{
			Path:         "/user_nu",
			Method:       http.MethodPost,
			Handler:      authHandlers.UpdateUserName,
			AuthRequired: true,
		},
		{
			Path:         "/users",
			Method:       http.MethodGet,
			Handler:      authHandlers.GetUsers,
			AuthRequired: true,
		},
	}
}
