package handlers

import (
	"Lescatit/pb"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	crawlAddr string
)

func init() {
	flag.StringVar(&crawlAddr, "crawl_addr", "http://localhost:9000", "crawler urld address")
}

// TestGetURLData tests retrieving the content in the url address.
func TestGetURLData(t *testing.T) {
	//get token
	url := crawlAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "New Test User",
		"Email":    "testuser@email.com",
		"Password": "testuser",
	}
	jsonSignInByte, err := json.Marshal(jsonSignIn)
	assert.NoError(t, err)
	payload := strings.NewReader(string(jsonSignInByte))

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	signIn := new(pb.SignInResponse)
	err = json.Unmarshal(body, signIn)
	assert.NoError(t, err)
	assert.NotNil(t, signIn)
	assert.NotNil(t, signIn.GetToken())
	assert.Equal(t, jsonSignIn["Name"], signIn.User.GetName())
	assert.Equal(t, jsonSignIn["Email"], signIn.User.GetEmail())

	// get data of the url
	url = crawlAddr + "/url_data"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Url", "https://sozcu.com.tr/")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedURLData := new(pb.GetURLDataResponse)
	err = json.Unmarshal(body, getedURLData)
	assert.NoError(t, err)
	assert.NotNil(t, getedURLData)
	assert.NotEmpty(t, getedURLData.GetUrl())
	assert.NotEmpty(t, getedURLData.GetData())
}

// TestGetURLsData tests retrieving the content in the url addresses.
func TestGetURLsData(t *testing.T) {
	//get token
	url := crawlAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "New Test User",
		"Email":    "testuser@email.com",
		"Password": "testuser",
	}
	jsonSignInByte, err := json.Marshal(jsonSignIn)
	assert.NoError(t, err)
	payload := strings.NewReader(string(jsonSignInByte))

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	signIn := new(pb.SignInResponse)
	err = json.Unmarshal(body, signIn)
	assert.NoError(t, err)
	assert.NotNil(t, signIn)
	assert.NotNil(t, signIn.GetToken())
	assert.Equal(t, jsonSignIn["Name"], signIn.User.GetName())
	assert.Equal(t, jsonSignIn["Email"], signIn.User.GetEmail())

	// get data of the urls
	url = crawlAddr + "/urls_data"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Urls", "https://sozcu.com.tr/,https://www.haberler.com/")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedURLsData := new([]pb.GetURLDataResponse)
	err = json.Unmarshal(body, getedURLsData)
	assert.NoError(t, err)
	assert.NotNil(t, getedURLsData)
}

// TestCrawlURL tests crawl the url
func TestCrawlURL(t *testing.T) {
	//get token
	url := crawlAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "New Test User",
		"Email":    "testuser@email.com",
		"Password": "testuser",
	}
	jsonSignInByte, err := json.Marshal(jsonSignIn)
	assert.NoError(t, err)
	payload := strings.NewReader(string(jsonSignInByte))

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	signIn := new(pb.SignInResponse)
	err = json.Unmarshal(body, signIn)
	assert.NoError(t, err)
	assert.NotNil(t, signIn)
	assert.NotNil(t, signIn.GetToken())
	assert.Equal(t, jsonSignIn["Name"], signIn.User.GetName())
	assert.Equal(t, jsonSignIn["Email"], signIn.User.GetEmail())

	// crawl url
	url = crawlAddr + "/url_crawl"
	jsonCrawlURL := map[string]interface{}{
		"Url": "https://sozcu.com.tr/",
		"CrawlRequest": map[string]interface{}{
			"UserAgent":            "colly - https://github.com/gocolly/colly",
			"MaxDepth":             0,
			"AllowedDomains":       []string{},
			"DisallowedDomains":    []string{},
			"DisallowedUrlFilters": []string{},
			"UrlFlters":            []string{},
			"UrlRevisit":           false,
			"MaxBodySize":          0,
			"RobotsTxt":            true,
		},
	}
	jsonCrawlURLByte, err := json.Marshal(jsonCrawlURL)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonCrawlURLByte))
	request, err = http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	crawledURL := new(pb.CrawlURLResponse)
	err = json.Unmarshal(body, crawledURL)
	assert.NoError(t, err)
	assert.NotNil(t, crawledURL)

}

// TestCrawlURLs tests crawl the urls
func TestCrawlURLs(t *testing.T) {
	//get token
	url := crawlAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "New Test User",
		"Email":    "testuser@email.com",
		"Password": "testuser",
	}
	jsonSignInByte, err := json.Marshal(jsonSignIn)
	assert.NoError(t, err)
	payload := strings.NewReader(string(jsonSignInByte))

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	signIn := new(pb.SignInResponse)
	err = json.Unmarshal(body, signIn)
	assert.NoError(t, err)
	assert.NotNil(t, signIn)
	assert.NotNil(t, signIn.GetToken())
	assert.Equal(t, jsonSignIn["Name"], signIn.User.GetName())
	assert.Equal(t, jsonSignIn["Email"], signIn.User.GetEmail())

	// crawl urls
	url = crawlAddr + "/urls_crawl"
	jsonCrawlURLs := map[string]interface{}{
		"Url": []string{"https://sozcu.com.tr/", "https://www.haberler.com/"},
		"CrawlRequest": map[string]interface{}{
			"UserAgent":            "colly - https://github.com/gocolly/colly",
			"MaxDepth":             0,
			"AllowedDomains":       []string{},
			"DisallowedDomains":    []string{},
			"DisallowedUrlFilters": []string{},
			"UrlFlters":            []string{},
			"UrlRevisit":           false,
			"MaxBodySize":          0,
			"RobotsTxt":            true,
		},
	}
	jsonCrawlURLsByte, err := json.Marshal(jsonCrawlURLs)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonCrawlURLsByte))
	request, err = http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	crawledURLs := new([]pb.CrawlURLResponse)
	err = json.Unmarshal(body, crawledURLs)
	assert.NoError(t, err)
	assert.NotNil(t, crawledURLs)
}
