package handlers

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/mtnmunuklu/lescatit/pb"

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
	jsonSignIn := map[string]string{
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
	signUp := new(pb.User)
	err = json.Unmarshal(body, signUp)
	assert.NoError(t, err)
	assert.NotNil(t, signUp)
	assert.Equal(t, jsonSignIn["Name"], signUp.GetName())
	assert.Equal(t, jsonSignIn["Email"], signUp.GetEmail())
	assert.Equal(t, signUp.GetRole(), "user")
	assert.Greater(t, current, signUp.GetCreated())
	assert.Greater(t, current, signUp.GetUpdated())
}

// TestSignIn tests the user login process.
func TestSignIn(t *testing.T) {
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.GetToken())
	assert.Equal(t, jsonSignIn["Name"], signIn.User.GetName())
	assert.Equal(t, jsonSignIn["Email"], signIn.User.GetEmail())
}

// TestGetUser tests pull user by id.
func TestGetUser(t *testing.T) {
	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.User.GetId())
	assert.NotEmpty(t, signIn.GetToken())

	// get user
	url = authAddr + "/user"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Email", "testuser@email.com")
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
	assert.Equal(t, signIn.User.GetId(), getUserResponse.GetId())
	assert.Equal(t, signIn.User.GetName(), getUserResponse.GetName())
	assert.Equal(t, signIn.User.GetEmail(), getUserResponse.GetEmail())
	assert.Equal(t, signIn.User.GetCreated(), getUserResponse.GetCreated())
	assert.Equal(t, signIn.User.GetUpdated(), getUserResponse.GetUpdated())

}

// TestDeleteUser tests delete the user.
func TestDeleteUser(t *testing.T) {

	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.User.GetId())
	assert.NotEmpty(t, signIn.GetToken())

	// delete user
	url = authAddr + "/user"
	request, err = http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Email", "testuser@email.com")
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	deletedUser := new(pb.DeleteUserResponse)
	err = json.Unmarshal(body, deletedUser)
	assert.NoError(t, err)
	assert.NotNil(t, deletedUser)
}

// TestChangeUserRole tests change the user role.
func TestChangeUserRole(t *testing.T) {
	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.User.GetId())
	assert.NotEmpty(t, signIn.GetToken())

	// change user role
	url = authAddr + "/user"
	jsonUpdateUser := map[string]interface{}{
		"Email": "testuser@email.com",
		"Role":  "admin",
	}
	jsonUpdateUserByte, err := json.Marshal(jsonUpdateUser)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUpdateUserByte))

	request, err = http.NewRequest("PATCH", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	changedUser := new(pb.User)
	err = json.Unmarshal(body, changedUser)
	assert.NoError(t, err)
	assert.NotNil(t, changedUser)
}

// TestUpdateUserPasword tests update the user password.
func TestUpdateUserPasword(t *testing.T) {
	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.User.GetId())
	assert.NotEmpty(t, signIn.GetToken())

	// update user password
	url = authAddr + "/user_pu"
	jsonUpdateUser := map[string]interface{}{
		"Email":       "testuser@email.com",
		"Password":    "testuser",
		"NewPassword": "newtestuser",
	}
	jsonUpdateUserByte, err := json.Marshal(jsonUpdateUser)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUpdateUserByte))

	request, err = http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	updatedUser := new(pb.User)
	err = json.Unmarshal(body, updatedUser)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
}

// TestUpdateUserEmail tests update the user email.
func TestUpdateUserEmail(t *testing.T) {
	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.User.GetId())
	assert.NotEmpty(t, signIn.GetToken())

	// update user email
	url = authAddr + "/user_eu"
	jsonUpdateUser := map[string]interface{}{
		"Email":    "testuser@email.com",
		"NewEmail": "newtestuser@email.com",
		"Password": "testuser",
	}
	jsonUpdateUserByte, err := json.Marshal(jsonUpdateUser)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUpdateUserByte))

	request, err = http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	updatedUser := new(pb.User)
	err = json.Unmarshal(body, updatedUser)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
}

// TestUpdateUserName tests update the user email.
func TestUpdateUserName(t *testing.T) {
	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.User.GetId())
	assert.NotEmpty(t, signIn.GetToken())

	// update user name
	url = authAddr + "/user_nu"
	jsonUpdateUser := map[string]interface{}{
		"Email":    "testuser@email.com",
		"Name":     "newtestuser",
		"Password": "testuser",
	}
	jsonUpdateUserByte, err := json.Marshal(jsonUpdateUser)
	assert.NoError(t, err)
	payload = strings.NewReader(string(jsonUpdateUserByte))

	request, err = http.NewRequest("POST", url, payload)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	request.Header.Add("Content-Type", "application/json")
	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	updatedUser := new(pb.User)
	err = json.Unmarshal(body, updatedUser)
	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
}

// TestGetUsers test pull all users from database.
func TestGetUsers(t *testing.T) {
	// get token
	url := authAddr + "/signin"
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
	assert.NotEmpty(t, signIn.GetToken())

	// get all users
	url = authAddr + "/users"
	request, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, request)

	authorization := "Bearer " + signIn.GetToken()
	request.Header.Add("Authorization", authorization)
	response, err = client.Do(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)

	body, err = ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotNil(t, body)

	getedUsers := new([]pb.User)
	err = json.Unmarshal(body, getedUsers)
	assert.NoError(t, err)
	assert.NotNil(t, getedUsers)
}
