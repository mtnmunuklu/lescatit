package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/handlers"
)

// NewCatzeRoutes provides the routing process for categorizer.
func NewCatzeRoutes(catzeHandlers handlers.CatzeHandlers) []*Route {
	return []*Route{
		{
			Method: http.MethodPost,
			Path:   "/url_catze",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.CategorizeURL(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/urls_catze",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.CategorizeURLs(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPut,
			Path:   "/cmodel",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.GenerateClassificationModel(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/cmodel",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.GetClassificationModel(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/cmodel",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.UpdateClassificationModel(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodDelete,
			Path:   "/cmodel",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.DeleteClassificationModel(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/cmodels",
			Handler: func(c *fiber.Ctx) error {
				return catzeHandlers.ListClassificationModels(c)
			},
			AuthRequired: true,
		},
	}
}
