package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewCatRoutes provides the routing process for categorization.
func NewCatRoutes(catHandlers handlers.CatHandlers) []*Route {
	return []*Route{
		{
			Path:         "/category",
			Method:       http.MethodGet,
			Handler:      catHandlers.GetCategory,
			AuthRequired: true,
		},
		{
			Path:         "/category",
			Method:       http.MethodPost,
			Handler:      catHandlers.UpdateCategory,
			AuthRequired: true,
		},
		{
			Path:         "/report",
			Method:       http.MethodPost,
			Handler:      catHandlers.ReportMiscategorization,
			AuthRequired: true,
		},
		{
			Path:         "/urls",
			Method:       http.MethodDelete,
			Handler:      catHandlers.DeleteURLs,
			AuthRequired: true,
		},
		{
			Path:         "/urls",
			Method:       http.MethodGet,
			Handler:      catHandlers.GetURLs,
			AuthRequired: true,
		},
		{
			Path:         "/url",
			Method:       http.MethodPut,
			Handler:      catHandlers.AddURL,
			AuthRequired: true,
		},
		{
			Path:         "/url",
			Method:       http.MethodDelete,
			Handler:      catHandlers.DeleteURL,
			AuthRequired: true,
		},
	}
}
