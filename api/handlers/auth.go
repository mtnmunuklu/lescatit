package handlers

import (
	"Lescatit/api/util"
	"Lescatit/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// AuthHandlers is the interface of the authentication operation.
type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

// AHandlers provides a connection with authentication service over proto buffer.
type AHandlers struct {
	authSvcClient pb.AuthServiceClient
}

// NewAuthHandlers creates a new AuthHandlers instance.
func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &AHandlers{authSvcClient: authSvcClient}
}

// SignUp performs the user registration process.
func (h *AHandlers) SignUp(w http.ResponseWriter, r *http.Request) {
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
	signUpRequest := new(pb.SignUpRequest)
	err = json.Unmarshal(body, signUpRequest)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	signUp, err := h.authSvcClient.SignUp(r.Context(), signUpRequest)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusCreated, signUp)
}

// SignIn performs the user login process.
func (h *AHandlers) SignIn(w http.ResponseWriter, r *http.Request) {
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
	signInRequest := new(pb.SignInRequest)
	err = json.Unmarshal(body, signInRequest)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	signIn, err := h.authSvcClient.SignIn(r.Context(), signInRequest)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, signIn)
}

// UpdateUser performs update the user.
func (h *AHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := util.AuthRequestWithId(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
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
	updateUserRequest := new(pb.UpdateUserRequest)
	err = json.Unmarshal(body, updateUserRequest)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updateUserRequest.Id = tokenPayload.UserId
	updatedUser, err := h.authSvcClient.UpdateUser(r.Context(), updateUserRequest)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, updatedUser)
}

// GetUser performs return the user by id.
func (h *AHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := util.AuthRequestWithId(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUser, err := h.authSvcClient.GetUser(r.Context(), &pb.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, getedUser)
}

// GetUsers list all users.
func (h *AHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := h.authSvcClient.ListUsers(r.Context(), &pb.ListUsersRequest{})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var getedUsers []*pb.User
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		getedUsers = append(getedUsers, user)
	}
	util.WriteAsJson(w, http.StatusOK, getedUsers)
}

// DeleteUser performs delete the user.
func (h *AHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

	tokenPayload, err := util.AuthRequestWithId(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	deletedUser, err := h.authSvcClient.DeleteUser(r.Context(), &pb.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, deletedUser)
}
