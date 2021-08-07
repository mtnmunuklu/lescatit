package resthandlers

import (
	"CWS/api/restutil"
	"CWS/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type CatHandlers interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	ReportMiscategorization(w http.ResponseWriter, r *http.Request)
	AddUrls(w http.ResponseWriter, r *http.Request)
	AddUrl(w http.ResponseWriter, r *http.Request)
	DeleteUrls(w http.ResponseWriter, r *http.Request)
	DeleteUrl(w http.ResponseWriter, r *http.Request)
	GetUrls(w http.ResponseWriter, r *http.Request)
}

type CHandlers struct {
	catSvcClient pb.CatServiceClient
}

func NewCatHandlers(catSvcClient pb.CatServiceClient) CatHandlers {
	return &CHandlers{catSvcClient: catSvcClient}
}

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

func (h *CHandlers) AddUrls(w http.ResponseWriter, r *http.Request) {
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
	urls := new(pb.AddUrlsRequest)
	err = json.Unmarshal(body, urls)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	stream, err := h.catSvcClient.AddUrls(r.Context(), urls)
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

func (h *CHandlers) AddUrl(w http.ResponseWriter, r *http.Request) {
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
	url := new(pb.AddUrlRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	addedUrl, err := h.catSvcClient.AddUrl(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, addedUrl)
}

func (h *CHandlers) DeleteUrls(w http.ResponseWriter, r *http.Request) {
	rUrls := strings.TrimSpace(r.Header.Get("Urls"))
	if rUrls == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	splittedUrls := strings.Split(rUrls, ",")
	urls := new(pb.DeleteUrlsRequest)
	urls.Urls = splittedUrls
	stream, err := h.catSvcClient.DeleteUrls(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var deletedUrls []*pb.DeleteUrlResponse
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

func (h *CHandlers) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	rUrl := strings.TrimSpace(r.Header.Get("Url"))
	if rUrl == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	url := new(pb.DeleteUrlRequest)
	url.Url = rUrl
	deletedUrl, err := h.catSvcClient.DeleteUrl(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", deletedUrl.Url)
	restutil.WriteAsJson(w, http.StatusNoContent, nil)
}

func (h *CHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
	rCategories := strings.TrimSpace(r.Header.Get("Categories"))
	rCount := strings.TrimSpace(r.Header.Get("Count"))
	if rCategories == "" || rCount == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	splittedCategories := strings.Split(rCategories, ",")
	urls := new(pb.ListUrlsRequest)
	urls.Categories = splittedCategories
	urls.Count = rCount
	stream, err := h.catSvcClient.ListUrls(r.Context(), urls)
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
