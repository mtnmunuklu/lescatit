package routes

import (
	"Lescatit/api/handlers"
	"net/http"
)

// NewCrawlRoutes provides the routing process for crawler.
func NewCrawlRoutes(crawlHandlers handlers.CrawlHandlers) []*Route {
	return []*Route{
		{
			Path:         "/api/url_data",
			Method:       http.MethodGet,
			Handler:      crawlHandlers.GetURLData,
			AuthRequired: true,
		},
		{
			Path:         "/api/urls_data",
			Method:       http.MethodGet,
			Handler:      crawlHandlers.GetURLsData,
			AuthRequired: true,
		},
		{
			Path:         "/api/url_crawl",
			Method:       http.MethodPost,
			Handler:      crawlHandlers.CrawlURL,
			AuthRequired: true,
		},
		{
			Path:         "/api/urls_crawl",
			Method:       http.MethodPost,
			Handler:      crawlHandlers.CrawlURLs,
			AuthRequired: true,
		},
	}
}
