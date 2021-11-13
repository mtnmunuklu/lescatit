package handlers

import (
	"Lescatit/api/util"
	"Lescatit/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// CatzeHandlers is the interface of the categorize operation.
type CatzeHandlers interface {
	CategorizeURL(w http.ResponseWriter, r *http.Request)
	CategorizeURLs(w http.ResponseWriter, r *http.Request)
	GenerateClassificationModel(w http.ResponseWriter, r *http.Request)
	GetClassificationModel(w http.ResponseWriter, r *http.Request)
	UpdateClassificationModel(w http.ResponseWriter, r *http.Request)
	DeleteClassificationModel(w http.ResponseWriter, r *http.Request)
	DeleteClassificationModels(w http.ResponseWriter, r *http.Request)
	ListClassificationModels(w http.ResponseWriter, r *http.Request)
}

// CzHandlers provides a connection with categorization service over proto buffer.
type CzHandlers struct {
	authSvcClient  pb.AuthServiceClient
	catzeSvcClient pb.CatzeServiceClient
}

// NewCatzeHandlers creates a new CatzeHandlers instance.
func NewCatzeHandlers(authSvcClient pb.AuthServiceClient, catzeSvcClient pb.CatzeServiceClient) CatzeHandlers {
	return &CzHandlers{authSvcClient: authSvcClient, catzeSvcClient: catzeSvcClient}
}

// CategorizeURL performs categorize the url.
func (h *CzHandlers) CategorizeURL(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	url := new(pb.CategorizeURLRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	categorizedURL, err := h.catzeSvcClient.CategorizeURL(r.Context(), url)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, categorizedURL)
}

// CategorizeURLs performs categorize the urls.
func (h *CzHandlers) CategorizeURLs(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	urls := new(pb.CategorizeURLsRequest)
	err = json.Unmarshal(body, urls)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	stream, err := h.catzeSvcClient.CategorizeURLs(r.Context(), urls)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var categorizedURLs []*pb.CategorizeURLResponse
	for {
		categorizedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		categorizedURLs = append(categorizedURLs, categorizedURL)
	}
	util.WriteAsJson(w, http.StatusOK, categorizedURLs)
}

// GenerateClassificationModel performs generate a classification model.
func (h *CzHandlers) GenerateClassificationModel(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// generate classification model
	if r.Body == nil {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	model := new(pb.GenerateClassificationModelRequest)
	err = json.Unmarshal(body, model)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	generatedModel, err := h.catzeSvcClient.GenerateClassificationModel(r.Context(), model)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, generatedModel)
}

// GetClassificationModel performs return the classification model.
func (h *CzHandlers) GetClassificationModel(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("Name")
	cmodel := new(pb.GetClassificationModelRequest)
	cmodel.Name = name
	getedCModel, err := h.catzeSvcClient.GetClassificationModel(r.Context(), cmodel)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, getedCModel)
}

// UpdateClassificationModel performs update the classification model.
func (h *CzHandlers) UpdateClassificationModel(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// update classification model
	if r.Body == nil {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	cmodel := new(pb.UpdateClassificationModelRequest)
	err = json.Unmarshal(body, cmodel)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedCModel, err := h.catzeSvcClient.UpdateClassificationModel(r.Context(), cmodel)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, updatedCModel)
}

// DeleteClassificationModel performs delete the classification model.
func (h *CzHandlers) DeleteClassificationModel(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// delete classification model
	name := r.Header.Get("Name")
	cmodel := new(pb.DeleteClassificationModelRequest)
	cmodel.Name = name
	deletedCModel, err := h.catzeSvcClient.DeleteClassificationModel(r.Context(), cmodel)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, deletedCModel)
}

// DeleteClassificationModels performs delete the classification models.
func (h *CzHandlers) DeleteClassificationModels(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// delete classification models
	names := r.Header.Get("Names")
	splittedNames := strings.Split(names, ",")
	cmodels := new(pb.DeleteClassificationModelsRequest)
	cmodels.Names = splittedNames
	stream, err := h.catzeSvcClient.DeleteClassificationModels(r.Context(), cmodels)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var deletedCModels []*pb.DeleteClassificationModelResponse
	for {
		deletedCModel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		deletedCModels = append(deletedCModels, deletedCModel)
	}
	util.WriteAsJson(w, http.StatusOK, deletedCModels)
}

// ListClassificationModels performs list all classification models.
func (h *CzHandlers) ListClassificationModels(w http.ResponseWriter, r *http.Request) {
	categories := r.Header.Get("Categories")
	count := r.Header.Get("Count")
	splittedCategories := strings.Split(categories, ",")
	cmodels := new(pb.ListClassificationModelsRequest)
	cmodels.Categories = splittedCategories
	cmodels.Count = count
	stream, err := h.catzeSvcClient.ListClassificationModels(r.Context(), cmodels)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var getedCModels []*pb.Classifier
	for {
		getedCModel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		getedCModels = append(getedCModels, getedCModel)
	}
	util.WriteAsJson(w, http.StatusOK, getedCModels)
}
