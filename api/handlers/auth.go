package handlers

import (
	"Lescatit/api/util"
	"Lescatit/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// AuthHandlers is the interface of the authentication operation.
type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
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
	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user.Created = time.Now().Unix()
	user.Updated = time.Now().Unix()
	user.Id = bson.NewObjectId().Hex()
	resp, err := h.authSvcClient.SignUp(r.Context(), user)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusCreated, resp)
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
	input := new(pb.SignInRequest)
	err = json.Unmarshal(body, input)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.authSvcClient.SignIn(r.Context(), input)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, resp)
}

// PutUser performs update the user.
func (h *AHandlers) PutUser(w http.ResponseWriter, r *http.Request) {
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
	user := new(pb.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	user.Id = tokenPayload.UserId
	resp, err := h.authSvcClient.UpdateUser(r.Context(), user)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, resp)
}

// GetUser performs return the user by id.
func (h *AHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	tokenPayload, err := util.AuthRequestWithId(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.authSvcClient.GetUser(r.Context(), &pb.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, resp)
}

// GetUsers list all users.
func (h *AHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	stream, err := h.authSvcClient.ListUsers(r.Context(), &pb.ListUsersRequest{})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	var users []*pb.User
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		users = append(users, user)
	}
	util.WriteAsJson(w, http.StatusOK, users)
}

// DeleteUser performs delete the user.
func (h *AHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

	tokenPayload, err := util.AuthRequestWithId(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	resp, err := h.authSvcClient.DeleteUser(r.Context(), &pb.GetUserRequest{Id: tokenPayload.UserId})
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", resp.Id)
	util.WriteAsJson(w, http.StatusNoContent, nil)
}
