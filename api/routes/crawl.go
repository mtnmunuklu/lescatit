package routes

import (
	"Lescatit/api/resthandlers"
	"net/http"
)

// NewCrawlRoutes provides the routing process for crawl.
func NewCrawlRoutes(crawlHandlers resthandlers.CrawlHandlers) []*Route {
	return []*Route{
		{
			Path:         "/url_data",
			Method:       http.MethodGet,
			Handler:      crawlHandlers.GetURLData,
			AuthRequired: true,
		},
		{
			Path:         "/urls_data",
			Method:       http.MethodGet,
			Handler:      crawlHandlers.GetURLsData,
			AuthRequired: true,
		},
		{
			Path:         "/url_crawl",
			Method:       http.MethodPost,
			Handler:      crawlHandlers.CrawlURL,
			AuthRequired: true,
		},
		{
			Path:         "/urls_crawl",
			Method:       http.MethodPost,
			Handler:      crawlHandlers.CrawlURLs,
			AuthRequired: true,
		},
	}
}
