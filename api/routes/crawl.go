package routes

import (
	"net/http"

	"github.com/mtnmunuklu/lescatit/api/handlers"
)

// NewCrawlRoutes provides the routing process for crawler.
func NewCrawlRoutes(crawlHandlers handlers.CrawlHandlers) []*Route {
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
