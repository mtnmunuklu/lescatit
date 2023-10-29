package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/util"
	"github.com/mtnmunuklu/lescatit/pb"
)

// CrawlHandlers is the interface of the crawler operation.
type CrawlHandlers interface {
	GetURLData(c *fiber.Ctx) error
	GetURLsData(c *fiber.Ctx) error
	CrawlURL(c *fiber.Ctx) error
	CrawlURLs(c *fiber.Ctx) error
}

// crawlHandlers provides a connection with crawler service over proto buffer.
type crawlHandlers struct {
	crawlSvcClient pb.CrawlServiceClient
}

// NewCrawlHandlers creates a new CrawlHandlers instance.
func NewCrawlHandlers(crawlSvcClient pb.CrawlServiceClient) CrawlHandlers {
	return &crawlHandlers{crawlSvcClient: crawlSvcClient}
}

// GetURLData provides to get the content in the url address.
func (h *crawlHandlers) GetURLData(c *fiber.Ctx) error {
	url := c.Get("Url")

	getURLDataRequest := &pb.GetURLDataRequest{Url: url}

	getedURLData, err := h.crawlSvcClient.GetURLData(c.Context(), getURLDataRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedURLData)
}

// GetURLsData provides to get the content in the url addresses.
func (h *crawlHandlers) GetURLsData(c *fiber.Ctx) error {
	urls := strings.Split(c.Get("Urls"), ",")
	types := strings.Split(c.Get("Types"), ",")

	getURLsDataRequest := new(pb.GetURLsDataRequest)
	for index, url := range urls {
		getURLDataRequest := new(pb.GetURLDataRequest)
		getURLDataRequest.Url = url
		if len(types) > index {
			getURLDataRequest.Type = types[index]
		}
		getURLsDataRequest.GetURLsDataRequest = append(getURLsDataRequest.GetURLsDataRequest, getURLDataRequest)
	}

	stream, err := h.crawlSvcClient.GetURLsData(c.Context(), getURLsDataRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var getedURLsData []*pb.GetURLDataResponse
	for {
		getedURLData, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusUnprocessableEntity, err)
		}

		getedURLsData = append(getedURLsData, getedURLData)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedURLsData)
}

// CrawlURL performs crawl the url.
func (h *crawlHandlers) CrawlURL(c *fiber.Ctx) error {
	crawlURLRequest := new(pb.CrawlURLRequest)
	if err := c.BodyParser(crawlURLRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	crawledURL, err := h.crawlSvcClient.CrawlURL(c.Context(), crawlURLRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, crawledURL)
}

// CrawlURLs performs crawl the urls.
func (h *crawlHandlers) CrawlURLs(c *fiber.Ctx) error {
	crawlURLsRequest := new(pb.CrawlURLsRequest)
	if err := c.BodyParser(crawlURLsRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	stream, err := h.crawlSvcClient.CrawlURLs(c.Context(), crawlURLsRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var crawledURLs []*pb.CrawlURLResponse
	for {
		crawledURL, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		crawledURLs = append(crawledURLs, crawledURL)
	}

	return util.WriteAsJSON(c, http.StatusOK, crawledURLs)
}
