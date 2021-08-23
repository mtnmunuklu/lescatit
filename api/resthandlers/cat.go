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
	fetchedUrl, err := h.catSvcClient.GetCategory(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, fetchedUrl)
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
	updatedUrl, err := h.catSvcClient.UpdateCategory(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, updatedUrl)
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
	reportedUrl, err := h.catSvcClient.ReportMiscategorization(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, reportedUrl)
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
	var addedUrls []*pb.Category
	for {
		addedUrl, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		addedUrls = append(addedUrls, addedUrl)
	}
	restutil.WriteAsJson(w, http.StatusOK, addedUrls)
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
	addedUrl, err := h.catSvcClient.AddURL(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, addedUrl)
}

// DeleteURLs performs delete the urls.
func (h *CHandlers) DeleteURLs(w http.ResponseWriter, r *http.Request) {
	rUrls := strings.TrimSpace(r.Header.Get("Urls"))
	if rUrls == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	splittedUrls := strings.Split(rUrls, ",")
	urls := new(pb.DeleteURLsRequest)
	urls.Urls = splittedUrls
	stream, err := h.catSvcClient.DeleteURLs(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var deletedUrls []*pb.DeleteURLResponse
	for {
		deletedUrl, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		deletedUrls = append(deletedUrls, deletedUrl)
	}
	restutil.WriteAsJson(w, http.StatusOK, deletedUrls)
}

// DeleteURL performs delete the url.
func (h *CHandlers) DeleteURL(w http.ResponseWriter, r *http.Request) {
	rUrl := strings.TrimSpace(r.Header.Get("Url"))
	if rUrl == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	url := new(pb.DeleteURLRequest)
	url.Url = rUrl
	deletedUrl, err := h.catSvcClient.DeleteURL(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", deletedUrl.Url)
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
	var fetchedUrls []*pb.Category
	for {
		fetchedUrl, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		fetchedUrls = append(fetchedUrls, fetchedUrl)
	}
	restutil.WriteAsJson(w, http.StatusOK, fetchedUrls)
}
