package handlers

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/mtnmunuklu/lescatit/pb"

	"github.com/stretchr/testify/assert"
)

var (
	catAddr string
)

func init() {
	flag.StringVar(&catAddr, "cat_addr", "http://localhost:9000", "categorization url address")
}

// TestGetCategory tests returning the category by url.
func TestGetCategory(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// get category
	url = catAddr + "/category"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Url", "https://www.examplect.com/")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedURL := new(pb.Category)
	err = json.Unmarshal(body, getedURL)
	assert.NoError(t, err)
	assert.NotNil(t, getedURL)
	assert.NotEmpty(t, getedURL.GetId())
	assert.NotEmpty(t, getedURL.GetUrl())
	assert.NotEmpty(t, getedURL.GetCategory())
	assert.NotEmpty(t, getedURL.GetData())
}

// TestUpdateCategory tests update the category.
func TestUpdateCategory(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// update category
	url = catzeAddr + "/category"
	jsonUURL := map[string]string{
		"Url":      "https://www.examplect.com/",
		"Category": "Gaming",
	}
	jsonUURLByte, err := json.Marshal(jsonUURL)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUURLByte))
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

	updatedURL := new(pb.Category)
	err = json.Unmarshal(body, updatedURL)
	assert.NoError(t, err)
	assert.NotNil(t, updatedURL)
	assert.NotEmpty(t, updatedURL.GetId())
	assert.NotEmpty(t, updatedURL.GetUrl())
	assert.NotEmpty(t, updatedURL.GetCategory())
	assert.NotEmpty(t, updatedURL.GetData())
	assert.Equal(t, jsonUURL["Url"], updatedURL.GetUrl())
	assert.Equal(t, jsonUURL["Category"], updatedURL.GetCategory())

}

// TestReportMiscategorization tests reporting miscategorization.
func TestReportMiscategorization(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// report miscategorization
	url = catzeAddr + "/url_report"
	jsonReportURL := map[string]string{
		"Url":    "https://www.examplect.com/",
		"Type":   "notnew",
		"Cmodel": "63d42860e2980ecbf7eb0d4d7ca9e488.nbc",
	}
	jsonReportURLByte, err := json.Marshal(jsonReportURL)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonReportURLByte))
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

	reportedURL := new(pb.Category)
	err = json.Unmarshal(body, reportedURL)
	assert.NoError(t, err)
	assert.NotNil(t, reportedURL)
	assert.NotEmpty(t, reportedURL.GetId())
	assert.NotEmpty(t, reportedURL.GetUrl())
	assert.NotEmpty(t, reportedURL.GetCategory())
	assert.NotEmpty(t, reportedURL.GetData())
	assert.Equal(t, jsonReportURL["Url"], reportedURL.GetUrl())
}

// TestAddURL tests adding the url.
func TestAddURL(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// add url
	url = catzeAddr + "/url"
	jsonAddURL := map[string]string{
		"Url":    "https://www.hurriyet.com.tr/",
		"Type":   "",
		"Cmodel": "63d42860e2980ecbf7eb0d4d7ca9e488.nbc",
	}
	jsonAddURLByte, err := json.Marshal(jsonAddURL)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonAddURLByte))
	request, err = http.NewRequest("PUT", url, payload)
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

	addedURL := new(pb.Category)
	err = json.Unmarshal(body, addedURL)
	assert.NoError(t, err)
	assert.NotNil(t, addedURL)
	assert.NotEmpty(t, addedURL.GetId())
	assert.NotEmpty(t, addedURL.GetUrl())
	assert.NotEmpty(t, addedURL.GetCategory())
	assert.NotEmpty(t, addedURL.GetData())
	assert.Equal(t, jsonAddURL["Url"], addedURL.GetUrl())
}

// TestDeleteURLs tests delete the urls.
func TestDeleteURLs(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// delete urls
	url = catAddr + "/urls"
	request, err = http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Urls", "https://www.hurriyet.com.tr/,https://www.example.com/")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deletedURLs := new([]pb.DeleteURLResponse)
	err = json.Unmarshal(body, deletedURLs)
	assert.NoError(t, err)
	assert.NotNil(t, deletedURLs)
}

// TestDeleteURL tests delete the urls.
func TestDeleteURL(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// delete url
	url = catAddr + "/url"
	request, err = http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Url", "https://www.examplect2.com/")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deletedURL := new(pb.DeleteURLResponse)
	err = json.Unmarshal(body, deletedURL)
	assert.NoError(t, err)
	assert.NotNil(t, deletedURL)
}

// TestGetURLs test listing the urls based on categories and count.
func TestGetURLs(t *testing.T) {
	// get token
	url := catAddr + "/signin"
	jsonSignIn := map[string]string{
		"Name":     "Test User",
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

	// get all urls
	url = catzeAddr + "/urls"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Categories", "News,Gaming")
	request.Header.Add("Count", "10")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedURLs := new([]pb.Category)
	err = json.Unmarshal(body, getedURLs)
	assert.NoError(t, err)
	assert.NotNil(t, getedURLs)
}
