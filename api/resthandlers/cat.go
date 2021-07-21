package resthandlers

import (
	"CWS/api/restutil"
	"CWS/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CatHandlers interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	ReportMiscategorization(w http.ResponseWriter, r *http.Request)
	GetUrls(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	AddUrl(w http.ResponseWriter, r *http.Request)
	DeleteUrl(w http.ResponseWriter, r *http.Request)
}

type catHandlers struct {
	catSvcClient pb.CatServiceClient
}

func NewCatHandlers(catSvcClient pb.CatServiceClient) CatHandlers {
	return &catHandlers{catSvcClient: catSvcClient}
}

func (h *catHandlers) GetCategory(w http.ResponseWriter, r *http.Request) {
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
	category := new(pb.GetCategoryRequest)
	err = json.Unmarshal(body, category)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.catSvcClient.GetCategory(r.Context(), category)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, resp)
}
func (h *catHandlers) ReportMiscategorization(w http.ResponseWriter, r *http.Request) {
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
	category := new(pb.GetCategoryRequest)
	err = json.Unmarshal(body, category)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.catSvcClient.ReportMiscategorization(r.Context(), category)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, resp)
}
func (h *catHandlers) GetUrls(w http.ResponseWriter, r *http.Request) {
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
	urls := new(pb.ListUrlsRequest)
	err = json.Unmarshal(body, urls)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	stream, err := h.catSvcClient.ListUrls(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var categories []*pb.Category
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		categories = append(categories, category)
	}
	restutil.WriteAsJson(w, http.StatusOK, categories)
}
func (h *catHandlers) UpdateCategory(w http.ResponseWriter, r *http.Request) {
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
	category := new(pb.UpdateCategoryRequest)
	err = json.Unmarshal(body, category)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.catSvcClient.UpdateCategory(r.Context(), category)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, resp)
}
func (h *catHandlers) AddUrl(w http.ResponseWriter, r *http.Request) {
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
	category := new(pb.Category)
	err = json.Unmarshal(body, category)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	category.Created = time.Now().Unix()
	category.Updated = time.Now().Unix()
	category.Id = bson.NewObjectId().Hex()
	category.Revision = 0
	resp, err := h.catSvcClient.AddUrl(r.Context(), category)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, resp)
}
func (h *catHandlers) DeleteUrl(w http.ResponseWriter, r *http.Request) {
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
	category := new(pb.GetCategoryRequest)
	err = json.Unmarshal(body, category)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.catSvcClient.DeleteUrl(r.Context(), category)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", resp.Url)
	restutil.WriteAsJson(w, http.StatusNoContent, nil)
}
