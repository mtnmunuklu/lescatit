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
	catzeSvcClient pb.CatzeServiceClient
}

// NewCatzeHandlers creates a new CatzeHandlers instance.
func NewCatzeHandlers(catzeSvcClient pb.CatzeServiceClient) CatzeHandlers {
	return &CzHandlers{catzeSvcClient: catzeSvcClient}
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

//DeleteClassificationModel performs generate a classification model
func (h *CzHandlers) GenerateClassificationModel(w http.ResponseWriter, r *http.Request) {
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
	createdModel, err := h.catzeSvcClient.GenerateClassificationModel(r.Context(), model)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, createdModel)
}

//DeleteClassificationModel performs return the classification model
func (h *CzHandlers) GetClassificationModel(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.Header.Get("Name"))
	if name == "" {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyHeader)
		return
	}
	cmodel := new(pb.GetClassificationModelRequest)
	cmodel.Name = name
	fetchedCModel, err := h.catzeSvcClient.GetClassificationModel(r.Context(), cmodel)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, fetchedCModel)
}

//DeleteClassificationModel performs update the classification model
func (h *CzHandlers) UpdateClassificationModel(w http.ResponseWriter, r *http.Request) {
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

//DeleteClassificationModel performs delete the classification model
func (h *CzHandlers) DeleteClassificationModel(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.Header.Get("Name"))
	if name == "" {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyHeader)
		return
	}
	cmodel := new(pb.DeleteClassificationModelRequest)
	cmodel.Name = name
	deletedCModel, err := h.catzeSvcClient.DeleteClassificationModel(r.Context(), cmodel)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, deletedCModel)
}

//DeleteClassificationModels performs delete the classification models
func (h *CzHandlers) DeleteClassificationModels(w http.ResponseWriter, r *http.Request) {
	names := strings.TrimSpace(r.Header.Get("Names"))
	if names == "" {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyHeader)
		return
	}
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

//DeleteClassificationModels performs list all classification models
func (h *CzHandlers) ListClassificationModels(w http.ResponseWriter, r *http.Request) {
	categories := strings.TrimSpace(r.Header.Get("Categories"))
	count := strings.TrimSpace(r.Header.Get("Count"))
	if categories == "" || count == "" {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyHeader)
		return
	}
	splittedCategories := strings.Split(categories, ",")
	cmodels := new(pb.ListClassificationModelsRequest)
	cmodels.Categories = splittedCategories
	cmodels.Count = count
	stream, err := h.catzeSvcClient.ListClassificationModels(r.Context(), cmodels)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var fetchedCModels []*pb.Classifier
	for {
		fetchedCModel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		fetchedCModels = append(fetchedCModels, fetchedCModel)
	}
	util.WriteAsJson(w, http.StatusOK, fetchedCModels)
}
