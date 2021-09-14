package routes

import (
	"Lescatit/api/resthandlers"
	"net/http"
)

// NewCatzeRoutes provides the routing process for categorize.
func NewCatzeRoutes(crawlHandlers resthandlers.CatzeHandlers) []*Route {
	return []*Route{
		{
			Path:         "/url_catze",
			Method:       http.MethodPost,
			Handler:      crawlHandlers.CategorizeURL,
			AuthRequired: true,
		},
		{
			Path:         "/urls_catze",
			Method:       http.MethodPost,
			Handler:      crawlHandlers.CategorizeURLs,
			AuthRequired: true,
		},
	}
}
