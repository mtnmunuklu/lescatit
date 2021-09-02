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

// CrawlHandlers is the interface of the crawler operation
type CrawlHandlers interface {
	GetURLData(w http.ResponseWriter, r *http.Request)
	GetURLsData(w http.ResponseWriter, r *http.Request)
	CrawlURL(w http.ResponseWriter, r *http.Request)
	CrawlURLs(w http.ResponseWriter, r *http.Request)
}

// CwlHandlers provides a connection with categorization service over proto buffer.
type CwlHandlers struct {
	crawlSvcClient pb.CrawlServiceClient
}

// NewCrawlHandlers creates a new CrawlHandlers instance.
func NewCrawlHandlers(crawlSvcClient pb.CrawlServiceClient) CrawlHandlers {
	return &CwlHandlers{crawlSvcClient: crawlSvcClient}
}

//GetURLData provides to get the content in the url address.
func (h *CwlHandlers) GetURLData(w http.ResponseWriter, r *http.Request) {
	rUrl := strings.TrimSpace(r.Header.Get("Url"))
	if rUrl == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	url := new(pb.GetURLDataRequest)
	url.Url = rUrl
	data, err := h.crawlSvcClient.GetURLData(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, data)
}

//GetURLsData provides to get the content in the url addresses.
func (h *CwlHandlers) GetURLsData(w http.ResponseWriter, r *http.Request) {
	rUrls := strings.TrimSpace(r.Header.Get("Urls"))
	if rUrls == "" {
		restutil.WriteError(w, http.StatusBadRequest, restutil.ErrEmptyHeader)
		return
	}
	splittedUrls := strings.Split(rUrls, ",")
	urls := new(pb.GetURLsDataRequest)
	urls.Urls = splittedUrls
	stream, err := h.crawlSvcClient.GetURLsData(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var data []*pb.GetURLDataResponse
	for {
		fetchedData, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		data = append(data, fetchedData)
	}
	restutil.WriteAsJson(w, http.StatusOK, data)
}

// CrawlURL performs crawl the url
func (h *CwlHandlers) CrawlURL(w http.ResponseWriter, r *http.Request) {
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
	url := new(pb.CrawlURLRequest)
	err = json.Unmarshal(body, url)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	crawledURL, err := h.crawlSvcClient.CrawlURL(r.Context(), url)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	restutil.WriteAsJson(w, http.StatusOK, crawledURL)
}

// CrawlURLs performs crawl the urls
func (h *CwlHandlers) CrawlURLs(w http.ResponseWriter, r *http.Request) {
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
	urls := new(pb.CrawlURLsRequest)
	err = json.Unmarshal(body, urls)
	if err != nil {
		restutil.WriteError(w, http.StatusBadRequest, err)
		return
	}
	stream, err := h.crawlSvcClient.CrawlURLs(r.Context(), urls)
	if err != nil {
		restutil.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var crawledURLs []*pb.CrawlURLResponse
	for {
		crawledURL, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			restutil.WriteError(w, http.StatusBadRequest, err)
			return
		}
		crawledURLs = append(crawledURLs, crawledURL)
	}
	restutil.WriteAsJson(w, http.StatusOK, crawledURLs)
}
