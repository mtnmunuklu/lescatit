package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/util"
	"github.com/mtnmunuklu/lescatit/pb"
)

// CatHandlers is the interface of the categorization operation.
type CatHandlers interface {
	GetCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	AddURL(c *fiber.Ctx) error
	DeleteURL(c *fiber.Ctx) error
	ReportMiscategorization(c *fiber.Ctx) error
	GetURLs(c *fiber.Ctx) error
}

// catHandlers provides a connection with categorization service over proto buffer.
type catHandlers struct {
	authSvcClient  pb.AuthServiceClient
	crawlSvcClient pb.CrawlServiceClient
	catzeSvcClient pb.CatzeServiceClient
	catSvcClient   pb.CatServiceClient
}

// NewCatHandlers creates a new CatHandlers instance.
func NewCatHandlers(authSvcClient pb.AuthServiceClient, crawlSvcClient pb.CrawlServiceClient, catzeSvcClient pb.CatzeServiceClient, catSvcClient pb.CatServiceClient) CatHandlers {
	return &catHandlers{authSvcClient: authSvcClient, crawlSvcClient: crawlSvcClient, catzeSvcClient: catzeSvcClient, catSvcClient: catSvcClient}
}

// GetCategory performs return the category by url.
func (h *catHandlers) GetCategory(c *fiber.Ctx) error {
	url := c.Get("Url")
	getCategoryRequest := &pb.GetCategoryRequest{Url: url}

	getedCategory, err := h.catSvcClient.GetCategory(c.Context(), getCategoryRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedCategory)
}

// UpdateCategory performs update the category.
func (h *catHandlers) UpdateCategory(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	updateCategoryRequest := new(pb.UpdateCategoryRequest)
	if err := c.BodyParser(updateCategoryRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	updatedCategory, err := h.catSvcClient.UpdateCategory(c.Context(), updateCategoryRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, updatedCategory)
}

// AddURL performs add the url.
func (h *catHandlers) AddURL(c *fiber.Ctx) error {

	addURLRequestCC := new(pb.AddURLRequestCC)
	if err := c.BodyParser(addURLRequestCC); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getURLDataRequest := &pb.GetURLDataRequest{Url: addURLRequestCC.Url, Type: addURLRequestCC.Type}
	getedURLData, err := h.crawlSvcClient.GetURLData(c.Context(), getURLDataRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	if getedURLData.GetStatus() != "" {
		categorizeURLRequest := &pb.CategorizeURLRequest{Url: getedURLData.Url, Data: getedURLData.Data, Cmodel: addURLRequestCC.Cmodel}
		categorizedURL, err := h.catzeSvcClient.CategorizeURL(c.Context(), categorizeURLRequest)
		if err != nil {
			return util.WriteError(c, http.StatusUnprocessableEntity, err)
		}

		getedData := new(pb.AddURLRequest)
		getedData.Url = getedURLData.GetUrl()
		getedData.Data = getedURLData.GetData()
		getedData.Status = getedURLData.GetStatus()
		getedData.Category = categorizedURL.GetCategory()
		addedURL, err := h.catSvcClient.AddURL(c.Context(), getedData)
		if err != nil {
			return util.WriteError(c, http.StatusUnprocessableEntity, err)
		}
		return util.WriteAsJSON(c, http.StatusOK, addedURL)
	}

	return util.WriteError(c, http.StatusUnprocessableEntity, err)
}

// DeleteURL performs delete the url.
func (h *catHandlers) DeleteURL(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	// delete url
	url := c.Get("Url")
	deleteURLRequest := &pb.DeleteURLRequest{Url: url}

	deletedURL, err := h.catSvcClient.DeleteURL(c.Context(), deleteURLRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, deletedURL)
}

// ReportMiscategorization reports miscategorization.
func (h *catHandlers) ReportMiscategorization(c *fiber.Ctx) error {
	reportURLRequestCC := new(pb.ReportURLRequestCC)
	if err := c.BodyParser(reportURLRequestCC); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	reportURLRequestCC.Type = "notnew"
	getURLDataRequest := &pb.GetURLDataRequest{Url: reportURLRequestCC.Url, Type: reportURLRequestCC.Type}
	getedURLData, err := h.crawlSvcClient.GetURLData(c.Context(), getURLDataRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	categorizeURLRequest := &pb.CategorizeURLRequest{Url: getedURLData.GetUrl(), Data: getedURLData.GetData(), Cmodel: reportURLRequestCC.Cmodel}
	categorizedURL, err := h.catzeSvcClient.CategorizeURL(c.Context(), categorizeURLRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	getedData := new(pb.ReportMiscategorizationRequest)
	getedData.Url = getedURLData.GetUrl()
	getedData.Data = getedURLData.GetData()
	getedData.Category = categorizedURL.GetCategory()

	reportedURL, err := h.catSvcClient.ReportMiscategorization(c.Context(), getedData)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, reportedURL)
}

// GetURLs performs list the urls based on categories and count.
func (h *catHandlers) GetURLs(c *fiber.Ctx) error {
	categories := strings.Split(c.Get("Categories"), ",")
	count := c.Get("Count")

	listURLsRequest := &pb.ListURLsRequest{Categories: categories, Count: count}

	stream, err := h.catSvcClient.ListURLs(c.Context(), listURLsRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var getedURLs []*pb.Category
	for {
		getedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		getedURLs = append(getedURLs, getedURL)
	}
	return util.WriteAsJSON(c, http.StatusOK, getedURLs)
}
