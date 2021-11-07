package handlers

import (
	"Lescatit/api/util"
	"Lescatit/pb"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// AuthHandlers is the interface of the authentication operation.
type AuthHandlers interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	ChangeUserRole(w http.ResponseWriter, r *http.Request)
	UpdateName(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	UpdateEmail(w http.ResponseWriter, r *http.Request)
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

// GetUser performs return the user by id.
func (h *AHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// get user
	email := strings.TrimSpace(r.Header.Get("Email"))
	if email == "" {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyHeader)
		return
	}
	user := new(pb.GetUserRequest)
	user.Email = email
	getedUser, err := h.authSvcClient.GetUser(r.Context(), user)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, getedUser)
}

// GetUsers list all users.
func (h *AHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// get users
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
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// delete user
	email := strings.TrimSpace(r.Header.Get("Email"))
	if email == "" {
		util.WriteError(w, http.StatusBadRequest, util.ErrEmptyHeader)
		return
	}
	user := new(pb.GetUserRequest)
	user.Email = email
	deletedUser, err := h.authSvcClient.DeleteUser(r.Context(), user)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, deletedUser)
}

// ChangeUserRole performs change the user role.
func (h *AHandlers) ChangeUserRole(w http.ResponseWriter, r *http.Request) {
	// check user role
	userId, err := util.GetUserIdFromToken(r)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	getedUserRole, err := h.authSvcClient.GetUserRole(r.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		util.WriteError(w, http.StatusUnauthorized, util.ErrUnauthorized)
		return
	}
	// change user role
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
	user := new(pb.ChangeUserRoleRequest)
	err = json.Unmarshal(body, user)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	changedUser, err := h.authSvcClient.ChangeUserRole(r.Context(), user)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, changedUser)
}

// UpdateUser performs update the user name.
func (h *AHandlers) UpdateName(w http.ResponseWriter, r *http.Request) {
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
	updateNameRequest := new(pb.UpdateNameRequest)
	err = json.Unmarshal(body, updateNameRequest)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedUser, err := h.authSvcClient.UpdateName(r.Context(), updateNameRequest)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, updatedUser)
}

// UpdateUser performs update the user password.
func (h *AHandlers) UpdatePassword(w http.ResponseWriter, r *http.Request) {
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
	updatePasswordRequest := new(pb.UpdatePasswordRequest)
	err = json.Unmarshal(body, updatePasswordRequest)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedUser, err := h.authSvcClient.UpdatePassword(r.Context(), updatePasswordRequest)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, updatedUser)
}

// UpdateUser performs update the user email.
func (h *AHandlers) UpdateEmail(w http.ResponseWriter, r *http.Request) {
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
	updateEmailRequest := new(pb.UpdateEmailRequest)
	err = json.Unmarshal(body, updateEmailRequest)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedUser, err := h.authSvcClient.UpdateEmail(r.Context(), updateEmailRequest)
	if err != nil {
		util.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	util.WriteAsJson(w, http.StatusOK, updatedUser)
}
