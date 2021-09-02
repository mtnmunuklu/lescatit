package resthandlers

import (
	"Lescatit/api/restutil"
	"Lescatit/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// CatHandlers is the interface of the categorization operation.
type CatHandlers interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	ReportMiscategorization(w http.ResponseWriter, r *http.Request)
	AddURLs(w http.ResponseWriter, r *http.Request)
	AddURL(w http.ResponseWriter, r *http.Request)
	DeleteURLs(w http.ResponseWriter, r *http.Request)
	DeleteURL(w http.ResponseWriter, r *http.Request)
	GetURLs(w http.ResponseWriter, r *http.Request)
}

// CHandlers provides a connection with categorization service over proto buffer.
type CHandlers struct {
	catSvcClient pb.CatServiceClient
}

// NewCatHandlers creates a new CatHandlers instance.
func NewCatHandlers(catSvcClient pb.CatServiceClient) CatHandlers {
	return &CHandlers{catSvcClient: catSvcClient}
}

// GetCategory performs return the category by url.
func (h *CHandlers) GetCategory(w http.ResponseWriter, r *http.Request) {
	rUrl := strings.TrimSpace(r.Header.Get("Url"))
	if rUrl == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	url := new(pb.GetCategoryRequest)
	url.Url = rUrl
	fetchedURL, err := h.catSvcClient.GetCategory(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, fetchedURL)
}

// UpdateCategory performs update the category.
func (h *CHandlers) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	url := new(pb.UpdateCategoryRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedURL, err := h.catSvcClient.UpdateCategory(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, updatedURL)
}

// ReportMiscategorization reports miscategorization.
func (h *CHandlers) ReportMiscategorization(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	url := new(pb.GetCategoryRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	reportedURL, err := h.catSvcClient.ReportMiscategorization(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, reportedURL)
}

// AddURLs performs add the urls.
func (h *CHandlers) AddURLs(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	urls := new(pb.AddURLsRequest)
	err = json.Unmarshal(body, urls)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	stream, err := h.catSvcClient.AddURLs(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var addedURLs []*pb.Category
	for {
		addedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		addedURLs = append(addedURLs, addedURL)
	}
	restutil.WriteAsJson(w, http.StatusOK, addedURLs)
}

// AddURL performs add the url.
func (h *CHandlers) AddURL(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyBody)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	url := new(pb.AddURLRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	addedURL, err := h.catSvcClient.AddURL(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, addedURL)
}

// DeleteURLs performs delete the urls.
func (h *CHandlers) DeleteURLs(w http.ResponseWriter, r *http.Request) {
	rURLs := strings.TrimSpace(r.Header.Get("Urls"))
	if rURLs == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	splittedURLs := strings.Split(rURLs, ",")
	urls := new(pb.DeleteURLsRequest)
	urls.Urls = splittedURLs
	stream, err := h.catSvcClient.DeleteURLs(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var deletedURLs []*pb.DeleteURLResponse
	for {
		deletedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		deletedURLs = append(deletedURLs, deletedURL)
	}
	restutil.WriteAsJson(w, http.StatusOK, deletedURLs)
}

// DeleteURL performs delete the url.
func (h *CHandlers) DeleteURL(w http.ResponseWriter, r *http.Request) {
	rURL := strings.TrimSpace(r.Header.Get("Url"))
	if rURL == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	url := new(pb.DeleteURLRequest)
	url.Url = rURL
	deletedURL, err := h.catSvcClient.DeleteURL(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", deletedURL.Url)
	restutil.WriteAsJson(w, http.StatusNoContent, nil)
}

// GetURLs performs list the urls based on categories and count.
func (h *CHandlers) GetURLs(w http.ResponseWriter, r *http.Request) {
	rCategories := strings.TrimSpace(r.Header.Get("Categories"))
	rCount := strings.TrimSpace(r.Header.Get("Count"))
	if rCategories == "" || rCount == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	splittedCategories := strings.Split(rCategories, ",")
	urls := new(pb.ListURLsRequest)
	urls.Categories = splittedCategories
	urls.Count = rCount
	stream, err := h.catSvcClient.ListURLs(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var fetchedURLs []*pb.Category
	for {
		fetchedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		fetchedURLs = append(fetchedURLs, fetchedURL)
	}
	restutil.WriteAsJson(w, http.StatusOK, fetchedURLs)
}
