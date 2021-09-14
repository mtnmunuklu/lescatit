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
	crawlSvcClient pb.CrawlServiceClient
	catzeSvcClient pb.CatzeServiceClient
	catSvcClient   pb.CatServiceClient
}

// NewCatHandlers creates a new CatHandlers instance.
func NewCatHandlers(crawlSvcClient pb.CrawlServiceClient, catzeSvcClient pb.CatzeServiceClient, catSvcClient pb.CatServiceClient) CatHandlers {
	return &CHandlers{crawlSvcClient: crawlSvcClient, catzeSvcClient: catzeSvcClient, catSvcClient: catSvcClient}
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
	url := new(pb.GetURLDataRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	url.Type = "notnew"
	fetchedURLData, err := h.crawlSvcClient.GetURLData(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	categorizedURL, err := h.catzeSvcClient.CategorizeURL(r.Context(),
		&pb.CategorizeURLRequest{Url: fetchedURLData.GetUrl(), Data: fetchedURLData.GetData()})
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	fetchedData := new(pb.ReportMiscategorizationRequest)
	fetchedData.Url = fetchedURLData.GetUrl()
	fetchedData.Data = fetchedURLData.GetData()
	fetchedData.Status = fetchedURLData.GetStatus()
	fetchedData.Category = categorizedURL.GetCategory()
	reportedURL, err := h.catSvcClient.ReportMiscategorization(r.Context(), fetchedData)
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
	urls := new(pb.GetURLsDataRequest)
	err = json.Unmarshal(body, urls)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// get urls data
	streamURLData, err := h.crawlSvcClient.GetURLsData(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	addURLsRequest := new(pb.AddURLsRequest)

	for {

		fetchedURLData, err := streamURLData.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}

		if fetchedURLData.GetStatus() != "" {
			categorizedURL, err := h.catzeSvcClient.CategorizeURL(r.Context(),
				&pb.CategorizeURLRequest{Url: fetchedURLData.Url, Data: fetchedURLData.Data})
			if err == nil {
				addURLRequest := new(pb.AddURLRequest)
				addURLRequest.Url = fetchedURLData.GetUrl()
				addURLRequest.Data = fetchedURLData.GetData()
				addURLRequest.Status = fetchedURLData.GetStatus()
				addURLRequest.Category = categorizedURL.GetCategory()
				addURLsRequest.AddURLsRequest = append(addURLsRequest.AddURLsRequest, addURLRequest)
			}
		}
	}

	// add urls
	stream, err := h.catSvcClient.AddURLs(r.Context(), addURLsRequest)
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
	url := new(pb.GetURLDataRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	fetchedURLData, err := h.crawlSvcClient.GetURLData(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	if fetchedURLData.GetStatus() != "" {
		categorizedURL, err := h.catzeSvcClient.CategorizeURL(r.Context(),
			&pb.CategorizeURLRequest{Url: fetchedURLData.Url, Data: fetchedURLData.Data})
		if err != nil {
			restutil.WriteError(w, http.StatusUnprocessableEntity, err)
			return
		}
		fetchedData := new(pb.AddURLRequest)
		fetchedData.Url = fetchedURLData.GetUrl()
		fetchedData.Data = fetchedURLData.GetData()
		fetchedData.Status = fetchedURLData.GetStatus()
		fetchedData.Category = categorizedURL.GetCategory()
		addedURL, err := h.catSvcClient.AddURL(r.Context(), fetchedData)
		if err != nil {
			restutil.WriteError(w, http.StatusUnprocessableEntity, err)
			return
		}
		restutil.WriteAsJson(w, http.StatusOK, addedURL)
	} else {
		restutil.WriteError(w, http.StatusUnprocessableEntity, restutil.ErrURLAlreadyExist)
		return
	}
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
