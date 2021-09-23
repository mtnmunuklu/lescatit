package handlers

import (
	"Lescatit/api/util"
	"Lescatit/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// CatzeHandlers is the interface of the categorize operation.
type CatzeHandlers interface {
	CategorizeURL(w http.ResponseWriter, r *http.Request)
	CategorizeURLs(w http.ResponseWriter, r *http.Request)
}

// CzHandlers provides a connection with categorization service over proto buffer.
type CzHandlers struct {
	catzeSvcClient pb.CatzeServiceClient
}

// NewCatzeHandlers creates a new CatzeHandlers instance.
func NewCatzeHandlers(catzeSvcClient pb.CatzeServiceClient) CatzeHandlers {
	return &CzHandlers{catzeSvcClient: catzeSvcClient}
}

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
