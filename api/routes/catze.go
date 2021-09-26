package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewCatzeRoutes provides the routing process for categorize.
func NewCatzeRoutes(catzeHandlers handlers.CatzeHandlers) []*Route {
	return []*Route{
		{
			Path:         "/url_catze",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.CategorizeURL,
			AuthRequired: true,
		},
		{
			Path:         "/urls_catze",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.CategorizeURLs,
			AuthRequired: true,
		},
		{
			Path:         "/cmodel_generate",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.GenerateClassificationModel,
			AuthRequired: true,
		},
	}
}
