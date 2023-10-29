package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/util"
	"github.com/mtnmunuklu/lescatit/pb"
)

// CatzeHandlers is the interface of the categorize operation.
type CatzeHandlers interface {
	CategorizeURL(c *fiber.Ctx) error
	CategorizeURLs(c *fiber.Ctx) error
	GenerateClassificationModel(c *fiber.Ctx) error
	GetClassificationModel(c *fiber.Ctx) error
	UpdateClassificationModel(c *fiber.Ctx) error
	DeleteClassificationModel(c *fiber.Ctx) error
	ListClassificationModels(c *fiber.Ctx) error
}

// catzeHandlers provides a connection with categorizer service over proto buffer.
type catzeHandlers struct {
	authSvcClient  pb.AuthServiceClient
	catzeSvcClient pb.CatzeServiceClient
}

// NewCatzeHandlers creates a new CatzeHandlers instance.
func NewCatzeHandlers(authSvcClient pb.AuthServiceClient, catzeSvcClient pb.CatzeServiceClient) CatzeHandlers {
	return &catzeHandlers{authSvcClient: authSvcClient, catzeSvcClient: catzeSvcClient}
}

// CategorizeURL performs categorize the url.
func (h *catzeHandlers) CategorizeURL(c *fiber.Ctx) error {
	categorizeURLRequest := new(pb.CategorizeURLRequest)
	if err := c.BodyParser(categorizeURLRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	categorizedURL, err := h.catzeSvcClient.CategorizeURL(c.Context(), categorizeURLRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, categorizedURL)
}

// CategorizeURLs performs categorize the urls.
func (h *catzeHandlers) CategorizeURLs(c *fiber.Ctx) error {
	categorizeURLsRequest := new(pb.CategorizeURLsRequest)
	if err := c.BodyParser(categorizeURLsRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	stream, err := h.catzeSvcClient.CategorizeURLs(c.Context(), categorizeURLsRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var categorizedURLs []*pb.CategorizeURLResponse
	for {
		categorizedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		categorizedURLs = append(categorizedURLs, categorizedURL)
	}

	return util.WriteAsJSON(c, http.StatusOK, categorizedURLs)
}

// GenerateClassificationModel performs generate a classification model.
func (h *catzeHandlers) GenerateClassificationModel(c *fiber.Ctx) error {
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

	generateClassificationModelRequest := new(pb.GenerateClassificationModelRequest)
	if err := c.BodyParser(generateClassificationModelRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	generatedClassificationModel, err := h.catzeSvcClient.GenerateClassificationModel(c.Context(), generateClassificationModelRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, generatedClassificationModel)
}

// GetClassificationModel performs return the classification model.
func (h *catzeHandlers) GetClassificationModel(c *fiber.Ctx) error {
	name := c.Get("Name")
	getClassificationModelRequest := &pb.GetClassificationModelRequest{Name: name}

	getedClassificationModel, err := h.catzeSvcClient.GetClassificationModel(c.Context(), getClassificationModelRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedClassificationModel)
}

// UpdateClassificationModel performs update the classification model.
func (h *catzeHandlers) UpdateClassificationModel(c *fiber.Ctx) error {
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

	updateClassificationModelRequest := new(pb.UpdateClassificationModelRequest)
	if err := c.BodyParser(updateClassificationModelRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	updatedClassificationModel, err := h.catzeSvcClient.UpdateClassificationModel(c.Context(), updateClassificationModelRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, updatedClassificationModel)
}

// DeleteClassificationModel performs delete the classification model.
func (h *catzeHandlers) DeleteClassificationModel(c *fiber.Ctx) error {
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

	// delete classification model
	name := c.Get("Name")
	deleteClassificationModelRequest := &pb.DeleteClassificationModelRequest{Name: name}

	deletedCModel, err := h.catzeSvcClient.DeleteClassificationModel(c.Context(), deleteClassificationModelRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, deletedCModel)
}

// ListClassificationModels performs list all classification models.
func (h *catzeHandlers) ListClassificationModels(c *fiber.Ctx) error {
	categories := strings.Split(c.Get("Categories"), ",")
	count := c.Get("Count")
	listClassificationModelsRequest := &pb.ListClassificationModelsRequest{Categories: categories, Count: count}

	stream, err := h.catzeSvcClient.ListClassificationModels(c.Context(), listClassificationModelsRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var getedClassificationModels []*pb.Classifier
	for {
		getedClassificationModel, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		getedClassificationModels = append(getedClassificationModels, getedClassificationModel)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedClassificationModels)
}
