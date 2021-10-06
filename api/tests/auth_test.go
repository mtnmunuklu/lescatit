package tests

import (
	"Lescatit/pb"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	authAddr string
)

func init() {
	flag.StringVar(&authAddr, "auth_addr", "http://localhost:9000", "authentication url address")

}

// TestSignUp tests the user registration process.
func TestSignUp(t *testing.T) {
	url := authAddr + "/signup"
	jsonSignIn := map[string]interface{}{
		"Name":     "Test User",
		"Email":    "testuser@email.com",
		"Password": "testuser",
	}
	jsonSignInByte, err := json.Marshal(jsonSignIn)
	assert.NoError(t, err)
	payload := strings.NewReader(string(jsonSignInByte))

	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	current := time.Now().Unix()
	signUpResponse := new(pb.User)
	err = json.Unmarshal(body, signUpResponse)
	assert.NoError(t, err)
	assert.NotNil(t, signUpResponse)
	assert.Equal(t, jsonSignIn["Name"], signUpResponse.GetName())
	assert.Equal(t, jsonSignIn["Email"], signUpResponse.GetEmail())
	assert.Greater(t, current, signUpResponse.GetCreated())
	assert.Greater(t, current, signUpResponse.GetUpdated())
}

// TestSignIn tests the user login process.
func TestSignIn(t *testing.T) {
	url := authAddr + "/signin"
	jsonSignIn := map[string]interface{}{
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

	signInResponse := new(pb.SignInResponse)
	err = json.Unmarshal(body, signInResponse)
	assert.NoError(t, err)
	assert.NotNil(t, signInResponse)
	assert.NotEmpty(t, signInResponse.GetToken())
	assert.Equal(t, jsonSignIn["Name"], signInResponse.User.GetName())
	assert.Equal(t, jsonSignIn["Email"], signInResponse.User.GetEmail())
}

// TestGetUser tests pull user by id.
func TestGetUser(t *testing.T) {
	// get token
	url := authAddr + "/signin"
	jsonSignIn := map[string]interface{}{
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

	signInResponse := new(pb.SignInResponse)
	err = json.Unmarshal(body, signInResponse)
	assert.NoError(t, err)
	assert.NotNil(t, signInResponse)
	assert.NotEmpty(t, signInResponse.User.GetId())
	assert.NotEmpty(t, signInResponse.GetToken())

	// get user
	url = authAddr + "/users/" + signInResponse.User.GetId()
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer" + signInResponse.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getUserResponse := new(pb.User)
	err = json.Unmarshal(body, getUserResponse)
	assert.NoError(t, err)
	assert.NotNil(t, getUserResponse)
	assert.Equal(t, signInResponse.User.GetId(), getUserResponse.GetId())
	assert.Equal(t, signInResponse.User.GetName(), getUserResponse.GetName())
	assert.Equal(t, signInResponse.User.GetEmail(), getUserResponse.GetEmail())
	assert.Equal(t, signInResponse.User.GetPassword(), getUserResponse.GetPassword())
	assert.Equal(t, signInResponse.User.GetCreated(), getUserResponse.GetCreated())
	assert.Equal(t, signInResponse.User.GetUpdated(), getUserResponse.GetUpdated())

}

// TestGetUsers test pull all users from database.
func TestGetUsers(t *testing.T) {
	// get token
	url := authAddr + "/signin"
	jsonSignIn := map[string]interface{}{
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

	signInResponse := new(pb.SignInResponse)
	err = json.Unmarshal(body, signInResponse)
	assert.NoError(t, err)
	assert.NotNil(t, signInResponse)
	assert.NotEmpty(t, signInResponse.GetToken())

	// get all users
	url = authAddr + "/users"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer" + signInResponse.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getUsersResponse := new([]pb.User)
	err = json.Unmarshal(body, getUsersResponse)
	assert.NoError(t, err)
	assert.NotNil(t, getUsersResponse)
}

// TestUpdateUser tests update the user.
func TestUpdateUser(t *testing.T) {
	// get token
	url := authAddr + "/signin"
	jsonSignIn := map[string]interface{}{
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

	signInResponse := new(pb.SignInResponse)
	err = json.Unmarshal(body, signInResponse)
	assert.NoError(t, err)
	assert.NotNil(t, signInResponse)
	assert.NotEmpty(t, signInResponse.User.GetId())
	assert.NotEmpty(t, signInResponse.GetToken())

	// update user
	url = authAddr + "/users/" + signInResponse.User.GetId()
	jsonUpdateUser := map[string]interface{}{
		"Name": "New Test User",
	}
	jsonUpdateUserByte, err := json.Marshal(jsonUpdateUser)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUpdateUserByte))

	request, err = http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	authorization := "Bearer" + signInResponse.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deleteUserResponse := new(pb.DeleteUserResponse)
	err = json.Unmarshal(body, deleteUserResponse)
	assert.NoError(t, err)
	assert.NotNil(t, deleteUserResponse)
}

// TestDeleteUser tests delete the user.
func TestDeleteUser(t *testing.T) {

	// get token
	url := authAddr + "/signin"
	jsonRaw := map[string]interface{}{
		"Name":     "New Test User",
		"Email":    "testuser@email.com",
		"Password": "testuser",
	}
	jsonRawByte, err := json.Marshal(jsonRaw)
	assert.NoError(t, err)
	payload := strings.NewReader(string(jsonRawByte))

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

	signInResponse := new(pb.SignInResponse)
	err = json.Unmarshal(body, signInResponse)
	assert.NoError(t, err)
	assert.NotNil(t, signInResponse)
	assert.NotEmpty(t, signInResponse.User.GetId())
	assert.NotEmpty(t, signInResponse.GetToken())

	// delete user
	url = authAddr + "/users/" + signInResponse.User.GetId()
	request, err = http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer" + signInResponse.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deleteUserResponse := new(pb.DeleteUserResponse)
	err = json.Unmarshal(body, deleteUserResponse)
	assert.NoError(t, err)
	assert.NotNil(t, deleteUserResponse)
}
