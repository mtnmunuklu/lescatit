package tests

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
	catzeAddr string
)

func init() {
	flag.StringVar(&catzeAddr, "api_addr", "http://localhost:9000", "categorizer url address")
}

// TestCategorizeURL tests to categorize the url.
func TestCategorizeURL(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// categorize url
	url = catzeAddr + "/url_catze"
	jsonCategorizeURL := map[string]string{
		"Url":    "https://sozcu.com.tr/",
		"Data":   "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wZWQgb3ZlciB0aGUgbGF6eSBkb2c=",
		"Cmodel": "d94087aec23fbbd167374555adfee686.nbc",
	}
	jsonCategorizeURLByte, err := json.Marshal(jsonCategorizeURL)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonCategorizeURLByte))
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

	categorizedURL := new(pb.CategorizeURLResponse)
	err = json.Unmarshal(body, categorizedURL)
	assert.NoError(t, err)
	assert.NotNil(t, categorizedURL)
	assert.Equal(t, jsonCategorizeURL["Url"], categorizedURL.GetUrl())
}

// TestCategorizeURLs tests to categorize the urls.
func TestCategorizeURLs(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// categorize urls
	url = catzeAddr + "/urls_catze"
	jsonCategorizeURLs := map[string]interface{}{
		"Urls": []interface{}{
			map[string]string{
				"Url":    "https://sozcu.com.tr/",
				"Data":   "ZGF0YQ==",
				"Cmodel": "d94087aec23fbbd167374555adfee686.nbc",
			},
			map[string]string{
				"Url":    "https://www.haberler.com/",
				"Data":   "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wZWQgb3ZlciB0aGUgbGF6eSBkb2c=",
				"Cmodel": "d94087aec23fbbd167374555adfee686.nbc",
			},
		},
	}
	jsonCategorizeURLsByte, err := json.Marshal(jsonCategorizeURLs)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonCategorizeURLsByte))
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

	categorizedURLs := new([]pb.CategorizeURLResponse)
	err = json.Unmarshal(body, categorizedURLs)
	assert.NoError(t, err)
	assert.NotNil(t, categorizedURLs)
}

// TestGenerateClassificationModel tests to generate a classification model.
func TestGenerateClassificationModel(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// generate classification model
	url = catzeAddr + "/cmodel"
	jsonGCModel := map[string]interface{}{
		"Category": "NB",
		"Urls": []interface{}{
			map[string]string{
				"Class": "New",
				"Data":  "ZGF0YQ==",
			},
			map[string]string{
				"Class": "Adult",
				"Data":  "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wZWQgb3ZlciB0aGUgbGF6eSBkb2c=",
			},
		},
	}
	jsonGCModelByte, err := json.Marshal(jsonGCModel)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonGCModelByte))
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

	generatedCModel := new(pb.Classifier)
	err = json.Unmarshal(body, generatedCModel)
	assert.NoError(t, err)
	assert.NotNil(t, generatedCModel)
	assert.NotEmpty(t, generatedCModel.GetId())
	assert.NotEmpty(t, generatedCModel.GetName())
	assert.NotEmpty(t, generatedCModel.GetCategory())
	assert.NotEmpty(t, generatedCModel.GetData())
}

// TestGetClassificationModel tests to return the classification model.
func TestGetClassificationModel(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// get classification model
	url = catzeAddr + "/cmodel"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Name", "3cbee7d6a2137387f86edee975f68e3f.nbc")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedCModel := new(pb.Classifier)
	err = json.Unmarshal(body, getedCModel)
	assert.NoError(t, err)
	assert.NotNil(t, getedCModel)
	assert.NotEmpty(t, getedCModel.GetId())
	assert.NotEmpty(t, getedCModel.GetName())
	assert.NotEmpty(t, getedCModel.GetCategory())
	assert.NotEmpty(t, getedCModel.GetData())
}

// TestUpdateClassificationModel tests update the classification model.
func TestUpdateClassificationModel(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// update classification model
	url = catzeAddr + "/cmodel"
	jsonUpdateCModel := map[string]string{
		"Name":     "9ebf5a938a7df5c198e792df0ff7dfd3.nbc",
		"Category": "KNN",
	}
	jsonUpdateCModelByte, err := json.Marshal(jsonUpdateCModel)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUpdateCModelByte))
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

	updatedCModel := new(pb.Classifier)
	err = json.Unmarshal(body, updatedCModel)
	assert.NoError(t, err)
	assert.NotNil(t, updatedCModel)
	assert.NotEmpty(t, updatedCModel.GetId())
	assert.NotEmpty(t, updatedCModel.GetName())
	assert.NotEmpty(t, updatedCModel.GetCategory())
	assert.NotEmpty(t, updatedCModel.GetData())
	assert.Equal(t, jsonUpdateCModel["Name"], updatedCModel.GetName())
	assert.Equal(t, jsonUpdateCModel["Category"], updatedCModel.GetCategory())

}

// TestDeleteClassificationModel tests to delete the classification model.
func TestDeleteClassificationModel(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// delete classification model
	url = catzeAddr + "/cmodel"
	request, err = http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Name", "3cbee7d6a2137387f86edee975f68e3f.nbc")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deletedCModel := new(pb.DeleteClassificationModelResponse)
	err = json.Unmarshal(body, deletedCModel)
	assert.NoError(t, err)
	assert.NotNil(t, deletedCModel)
}

// TestDeleteClassificationModels tests to delete the classification models.
func TestDeleteClassificationModels(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// delete classification models
	url = catzeAddr + "/cmodels"
	request, err = http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Name", "3cbee7d6a2137387f86edee975f68e3f.nbc,3cbee7d6a21334567f86edee975f68e3f.nbc")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deletedCModels := new([]pb.DeleteClassificationModelResponse)
	err = json.Unmarshal(body, deletedCModels)
	assert.NoError(t, err)
	assert.NotNil(t, deletedCModels)
}

// TestListClassificationModels tests listing all classification models.
func TestListClassificationModels(t *testing.T) {
	// get token
	url := catzeAddr + "/signin"
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

	// get all classification models
	url = catzeAddr + "/cmodels"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Categories", "NB,KNN")
	request.Header.Add("Count", "10")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedCModels := new([]pb.Classifier)
	err = json.Unmarshal(body, getedCModels)
	assert.NoError(t, err)
	assert.NotNil(t, getedCModels)
}
