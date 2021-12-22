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

// CatHandlers is the interface of the categorization operation.
type CatHandlers interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	AddURL(w http.ResponseWriter, r *http.Request)
	DeleteURL(w http.ResponseWriter, r *http.Request)
	ReportMiscategorization(w http.ResponseWriter, r *http.Request)
	DeleteURLs(w http.ResponseWriter, r *http.Request)
	GetURLs(w http.ResponseWriter, r *http.Request)
}

// CHandlers provides a connection with categorization service over proto buffer.
type CHandlers struct {
	authSvcClient  pb.AuthServiceClient
	crawlSvcClient pb.CrawlServiceClient
	catzeSvcClient pb.CatzeServiceClient
	catSvcClient   pb.CatServiceClient
}

// NewCatHandlers creates a new CatHandlers instance.
func NewCatHandlers(authSvcClient pb.AuthServiceClient, crawlSvcClient pb.CrawlServiceClient, catzeSvcClient pb.CatzeServiceClient, catSvcClient pb.CatServiceClient) CatHandlers {
	return &CHandlers{authSvcClient: authSvcClient, crawlSvcClient: crawlSvcClient, catzeSvcClient: catzeSvcClient, catSvcClient: catSvcClient}
}

// GetCategory performs return the category by url.
func (h *CHandlers) GetCategory(w http.ResponseWriter, r *http.Request) {
	rUrl := r.Header.Get("Url")
	url := new(pb.GetCategoryRequest)
	url.Url = rUrl

	getedURL, err := h.catSvcClient.GetCategory(r.Context(), url)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	util.WriteAsJson(w, http.StatusOK, getedURL)
}

// UpdateCategory performs update the category.
func (h *CHandlers) UpdateCategory(w http.ResponseWriter, r *http.Request) {
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

	// update category
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

	url := new(pb.UpdateCategoryRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	updatedURL, err := h.catSvcClient.UpdateCategory(r.Context(), url)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	util.WriteAsJson(w, http.StatusOK, updatedURL)
}

// AddURL performs add the url.
func (h *CHandlers) AddURL(w http.ResponseWriter, r *http.Request) {
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

	url := new(pb.AddURLRequestCC)
	err = json.Unmarshal(body, url)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	getedURLData, err := h.crawlSvcClient.GetURLData(r.Context(),
		&pb.GetURLDataRequest{Url: url.Url, Type: url.Type})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if getedURLData.GetStatus() != "" {
		categorizedURL, err := h.catzeSvcClient.CategorizeURL(r.Context(),
			&pb.CategorizeURLRequest{Url: getedURLData.Url, Data: getedURLData.Data, Cmodel: url.Cmodel})
		if err != nil {
			util.WriteError(w, http.StatusUnprocessableEntity, err)
			return
		}

		getedData := new(pb.AddURLRequest)
		getedData.Url = getedURLData.GetUrl()
		getedData.Data = getedURLData.GetData()
		getedData.Status = getedURLData.GetStatus()
		getedData.Category = categorizedURL.GetCategory()
		addedURL, err := h.catSvcClient.AddURL(r.Context(), getedData)
		if err != nil {
			util.WriteError(w, http.StatusUnprocessableEntity, err)
			return
		}

		util.WriteAsJson(w, http.StatusOK, addedURL)
	} else {
		util.WriteError(w, http.StatusUnprocessableEntity, util.ErrExistURL)
		return
	}
}

// DeleteURL performs delete the url.
func (h *CHandlers) DeleteURL(w http.ResponseWriter, r *http.Request) {
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

	// delete url
	rURL := r.Header.Get("Url")
	url := new(pb.DeleteURLRequest)
	url.Url = rURL

	deletedURL, err := h.catSvcClient.DeleteURL(r.Context(), url)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	util.WriteAsJson(w, http.StatusOK, deletedURL)
}

// ReportMiscategorization reports miscategorization.
func (h *CHandlers) ReportMiscategorization(w http.ResponseWriter, r *http.Request) {
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

	url := new(pb.AddURLRequestCC)
	err = json.Unmarshal(body, url)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	url.Type = "notnew"
	getedURLData, err := h.crawlSvcClient.GetURLData(r.Context(),
		&pb.GetURLDataRequest{Url: url.Url, Type: url.Type})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	categorizedURL, err := h.catzeSvcClient.CategorizeURL(r.Context(),
		&pb.CategorizeURLRequest{Url: getedURLData.GetUrl(), Data: getedURLData.GetData(), Cmodel: url.Cmodel})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	getedData := new(pb.ReportMiscategorizationRequest)
	getedData.Url = getedURLData.GetUrl()
	getedData.Data = getedURLData.GetData()
	getedData.Status = getedURLData.GetStatus()
	getedData.Category = categorizedURL.GetCategory()

	reportedURL, err := h.catSvcClient.ReportMiscategorization(r.Context(), getedData)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	util.WriteAsJson(w, http.StatusOK, reportedURL)
}

// DeleteURLs performs delete the urls.
func (h *CHandlers) DeleteURLs(w http.ResponseWriter, r *http.Request) {
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

	// delete urls
	rURLs := r.Header.Get("Urls")
	splittedURLs := strings.Split(rURLs, ",")
	urls := new(pb.DeleteURLsRequest)
	urls.Urls = splittedURLs

	stream, err := h.catSvcClient.DeleteURLs(r.Context(), urls)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var deletedURLs []*pb.DeleteURLResponse
	for {
		deletedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}

		deletedURLs = append(deletedURLs, deletedURL)
	}

	util.WriteAsJson(w, http.StatusOK, deletedURLs)
}

// GetURLs performs list the urls based on categories and count.
func (h *CHandlers) GetURLs(w http.ResponseWriter, r *http.Request) {
	rCategories := r.Header.Get("Categories")
	rCount := r.Header.Get("Count")
	splittedCategories := strings.Split(rCategories, ",")

	urls := new(pb.ListURLsRequest)
	urls.Categories = splittedCategories
	urls.Count = rCount

	stream, err := h.catSvcClient.ListURLs(r.Context(), urls)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var getedURLs []*pb.Category
	for {
		getedURL, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}

		getedURLs = append(getedURLs, getedURL)
	}

	util.WriteAsJson(w, http.StatusOK, getedURLs)
}
