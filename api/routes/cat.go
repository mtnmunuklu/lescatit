package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewCatRoutes provides the routing process for categorization.
func NewCatRoutes(catHandlers handlers.CatHandlers) []*Route {
	return []*Route{
		{
			Path:         "/api/category",
			Method:       http.MethodGet,
			Handler:      catHandlers.GetCategory,
			AuthRequired: true,
		},
		{
			Path:         "/api/category",
			Method:       http.MethodPost,
			Handler:      catHandlers.UpdateCategory,
			AuthRequired: true,
		},
		{
			Path:         "/api/url",
			Method:       http.MethodPut,
			Handler:      catHandlers.AddURL,
			AuthRequired: true,
		},
		{
			Path:         "/api/url",
			Method:       http.MethodDelete,
			Handler:      catHandlers.DeleteURL,
			AuthRequired: true,
		},
		{
			Path:         "/api/url_report",
			Method:       http.MethodPost,
			Handler:      catHandlers.ReportMiscategorization,
			AuthRequired: true,
		},
		{
			Path:         "/api/urls",
			Method:       http.MethodDelete,
			Handler:      catHandlers.DeleteURLs,
			AuthRequired: true,
		},
		{
			Path:         "/api/urls",
			Method:       http.MethodGet,
			Handler:      catHandlers.GetURLs,
			AuthRequired: true,
		},
	}
}
