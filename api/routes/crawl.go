package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/handlers"
)

// NewCrawlRoutes provides the routing process for crawler.
func NewCrawlRoutes(crawlHandlers handlers.CrawlHandlers) []*Route {
	return []*Route{
		{
			Method: http.MethodGet,
			Path:   "/url_data",
			Handler: func(c *fiber.Ctx) error {
				return crawlHandlers.GetURLData(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/urls_data",
			Handler: func(c *fiber.Ctx) error {
				return crawlHandlers.GetURLsData(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/url_crawl",
			Handler: func(c *fiber.Ctx) error {
				return crawlHandlers.CrawlURL(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/urls_crawl",
			Handler: func(c *fiber.Ctx) error {
				return crawlHandlers.CrawlURLs(c)
			},
			AuthRequired: true,
		},
	}
}
