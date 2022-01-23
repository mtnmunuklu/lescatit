package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewCatzeRoutes provides the routing process for categorizer.
func NewCatzeRoutes(catzeHandlers handlers.CatzeHandlers) []*Route {
	return []*Route{
		{
			Path:         "/api/url_catze",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.CategorizeURL,
			AuthRequired: true,
		},
		{
			Path:         "/api/urls_catze",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.CategorizeURLs,
			AuthRequired: true,
		},
		{
			Path:         "/api/cmodel",
			Method:       http.MethodPut,
			Handler:      catzeHandlers.GenerateClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/api/cmodel",
			Method:       http.MethodGet,
			Handler:      catzeHandlers.GetClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/api/cmodel",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.UpdateClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/api/cmodel",
			Method:       http.MethodDelete,
			Handler:      catzeHandlers.DeleteClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/api/cmodels",
			Method:       http.MethodDelete,
			Handler:      catzeHandlers.DeleteClassificationModels,
			AuthRequired: true,
		},
		{
			Path:         "/api/cmodels",
			Method:       http.MethodGet,
			Handler:      catzeHandlers.ListClassificationModels,
			AuthRequired: true,
		},
	}
}
