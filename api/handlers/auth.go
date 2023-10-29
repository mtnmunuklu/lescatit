package handlers

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/lescatit/api/util"
	"github.com/mtnmunuklu/lescatit/pb"
)

// AuthHandlers is the interface of the authentication operation.
type AuthHandlers interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	ChangeUserRole(c *fiber.Ctx) error
	UpdateUserPassword(c *fiber.Ctx) error
	UpdateUserEmail(c *fiber.Ctx) error
	UpdateUserName(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}

// authHandlers provides a connection with authentication service over proto buffer.
type authHandlers struct {
	authSvcClient pb.AuthServiceClient
}

// NewAuthHandlers creates a new AuthHandlers instance.
func NewAuthHandlers(authSvcClient pb.AuthServiceClient) AuthHandlers {
	return &authHandlers{authSvcClient: authSvcClient}
}

// SignUp performs the user registration process.
func (h *authHandlers) SignUp(c *fiber.Ctx) error {
	signUpRequest := new(pb.SignUpRequest)
	if err := c.BodyParser(signUpRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	signUp, err := h.authSvcClient.SignUp(c.Context(), signUpRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusCreated, signUp)
}

// SignIn performs the user login process.
func (h *authHandlers) SignIn(c *fiber.Ctx) error {
	signInRequest := new(pb.SignInRequest)
	if err := c.BodyParser(signInRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	signIn, err := h.authSvcClient.SignIn(c.Context(), signInRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, signIn)
}

// GetUser returns the user by email.
func (h *authHandlers) GetUser(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	email := c.Get("Email")
	user := &pb.GetUserRequest{Email: email}

	getedUser, err := h.authSvcClient.GetUser(c.Context(), user)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedUser)
}

// DeleteUser deletes the user by email.
func (h *authHandlers) DeleteUser(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	email := c.Get("Email")
	user := &pb.DeleteUserRequest{Email: email}

	deletedUser, err := h.authSvcClient.DeleteUser(c.Context(), user)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, deletedUser)
}

// ChangeUserRole changes the user role.
func (h *authHandlers) ChangeUserRole(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	changeUserRoleRequest := new(pb.ChangeUserRoleRequest)
	if err := c.BodyParser(changeUserRoleRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	changedUser, err := h.authSvcClient.ChangeUserRole(c.Context(), changeUserRoleRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, changedUser)
}

// UpdateUserPassword updates the user password.
func (h *authHandlers) UpdateUserPassword(c *fiber.Ctx) error {
	updatePasswordRequest := new(pb.UpdateUserPasswordRequest)
	if err := c.BodyParser(updatePasswordRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	updatedUser, err := h.authSvcClient.UpdateUserPassword(c.Context(), updatePasswordRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, updatedUser)
}

// UpdateUserEmail updates the user email.
func (h *authHandlers) UpdateUserEmail(c *fiber.Ctx) error {
	updateEmailRequest := new(pb.UpdateUserEmailRequest)
	if err := c.BodyParser(updateEmailRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	updatedUser, err := h.authSvcClient.UpdateUserEmail(c.Context(), updateEmailRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, updatedUser)
}

// UpdateUserName updates the user name.
func (h *authHandlers) UpdateUserName(c *fiber.Ctx) error {
	updateNameRequest := new(pb.UpdateUserNameRequest)
	if err := c.BodyParser(updateNameRequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	updatedUser, err := h.authSvcClient.UpdateUserName(c.Context(), updateNameRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, updatedUser)
}

// GetUsers lists all users.
func (h *authHandlers) GetUsers(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	stream, err := h.authSvcClient.ListUsers(c.Context(), &pb.ListUsersRequest{})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var getedUsers []*pb.User
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		getedUsers = append(getedUsers, user)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedUsers)
}
