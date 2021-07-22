package routes

import (
	"CWS/api/resthandlers"
	"net/http"
)

func NewCatRoutes(catHandlers resthandlers.CatHandlers) []*Route {
	return []*Route{
		{
			Path:         "/category",
			Method:       http.MethodGet,
			Handler:      catHandlers.GetCategory,
			AuthRequired: true,
		},
		{
			Path:         "/category",
			Method:       http.MethodPut,
			Handler:      catHandlers.UpdateCategory,
			AuthRequired: true,
		},
		{
			Path:         "/report",
			Method:       http.MethodPost,
			Handler:      catHandlers.ReportMiscategorization,
			AuthRequired: true,
		},
		{
			Path:         "/urls",
			Method:       http.MethodPost,
			Handler:      catHandlers.AddUrls,
			AuthRequired: true,
		},
		{
			Path:         "/urls",
			Method:       http.MethodDelete,
			Handler:      catHandlers.DeleteUrls,
			AuthRequired: true,
		},
		{
			Path:         "/urls",
			Method:       http.MethodGet,
			Handler:      catHandlers.GetUrls,
			AuthRequired: true,
		},
		{
			Path:         "/url",
			Method:       http.MethodPost,
			Handler:      catHandlers.AddUrl,
			AuthRequired: true,
		},
		{
			Path:         "/url",
			Method:       http.MethodDelete,
			Handler:      catHandlers.DeleteUrl,
			AuthRequired: true,
		},
	}
}
