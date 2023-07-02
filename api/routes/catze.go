package routes

import (
	"net/http"

	"github.com/mtnmunuklu/lescatit/api/handlers"
)

// NewCatzeRoutes provides the routing process for categorizer.
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
			Path:         "/cmodel",
			Method:       http.MethodPut,
			Handler:      catzeHandlers.GenerateClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/cmodel",
			Method:       http.MethodGet,
			Handler:      catzeHandlers.GetClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/cmodel",
			Method:       http.MethodPost,
			Handler:      catzeHandlers.UpdateClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/cmodel",
			Method:       http.MethodDelete,
			Handler:      catzeHandlers.DeleteClassificationModel,
			AuthRequired: true,
		},
		{
			Path:         "/cmodels",
			Method:       http.MethodDelete,
			Handler:      catzeHandlers.DeleteClassificationModels,
			AuthRequired: true,
		},
		{
			Path:         "/cmodels",
			Method:       http.MethodGet,
			Handler:      catzeHandlers.ListClassificationModels,
			AuthRequired: true,
		},
	}
}
