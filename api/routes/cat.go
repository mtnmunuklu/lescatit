package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/handlers"
)

// NewCatRoutes provides the routing process for categorization.
func NewCatRoutes(catHandlers handlers.CatHandlers) []*Route {
	return []*Route{
		{
			Method: http.MethodGet,
			Path:   "/category",
			Handler: func(c *fiber.Ctx) error {
				return catHandlers.GetCategory(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/category",
			Handler: func(c *fiber.Ctx) error {
				return catHandlers.UpdateCategory(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPut,
			Path:   "/url",
			Handler: func(c *fiber.Ctx) error {
				return catHandlers.AddURL(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodDelete,
			Path:   "/url",
			Handler: func(c *fiber.Ctx) error {
				return catHandlers.DeleteURL(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/url_report",
			Handler: func(c *fiber.Ctx) error {
				return catHandlers.ReportMiscategorization(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/urls",
			Handler: func(c *fiber.Ctx) error {
				return catHandlers.GetURLs(c)
			},
			AuthRequired: true,
		},
	}
}
