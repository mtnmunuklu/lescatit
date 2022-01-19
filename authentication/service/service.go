package service

import (
	"Lescatit/authentication/models"
	"Lescatit/authentication/repository"
	"Lescatit/authentication/util"
	"Lescatit/pb"
	"Lescatit/security"
	"context"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AuthService provides usersRepository for authentication service.
type AuthService struct {
	usersRepository repository.UsersRepository
}

// NewAuthService creates a new AuthService instance.
func NewAuthService(usersRepository repository.UsersRepository) pb.AuthServiceServer {
	return &AuthService{usersRepository: usersRepository}
}

// SignUp performs the user registration process.
func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.User, error) {
	err := util.ValidateSignUp(req)
	if err != nil {
		return nil, err
	}

	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}
	found, err := s.usersRepository.GetByEmail(normalizedEmail)

	if err == mgo.ErrNotFound {
		user := new(models.User)
		user.Id = bson.NewObjectId()
		user.Name = strings.TrimSpace(req.GetName())
		user.Email = normalizedEmail
		user.Password, err = security.EncryptPassword(req.GetPassword())
		if err != nil {
			return nil, err
		}
		if user.Email == "admin@lescatit.com" {
			user.Role = "admin"
		} else {
			user.Role = "user"
		}
		user.Created = time.Now()
		user.Updated = time.Now()
		err := s.usersRepository.Save(user)
		if err != nil {
			return nil, util.ErrCreateUser
		}
		return user.ToProtoBuffer(), nil
	}

	if found == nil {
		return nil, util.ErrNotPerformedOperation
	}

	return nil, util.ErrExistEmail
}

// SignIn performs the user login process.
func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}

	user, err := s.usersRepository.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}

	err = security.VerifyPassword(user.Password, req.GetPassword())
	if err != nil {
		return nil, util.ErrMismatchedPassword
	}

	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		return nil, util.ErrFailedSignIn
	}

	return &pb.SignInResponse{User: user.ToProtoBuffer(), Token: token}, nil
}

// GetUser performs return the user by id.
func (s *AuthService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}

	found, err := s.usersRepository.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}

	return found.ToProtoBuffer(), nil
}

// DeleteUser performs delete the user.
func (s *AuthService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	normalizedEmail := util.NormalizeEmail(req.Email)
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}

	user, err := s.usersRepository.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}

	err = s.usersRepository.DeleteByEmail(user.Email)
	if err != nil {
		return nil, util.ErrDeleteUser
	}

	return &pb.DeleteUserResponse{Email: user.Email}, nil
}

// ChangeUserRole performs change the user role.
func (s *AuthService) ChangeUserRole(ctx context.Context, req *pb.ChangeUserRoleRequest) (*pb.User, error) {
	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}

	user, err := s.usersRepository.GetByEmail(req.GetEmail())
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}

	req.Role = strings.TrimSpace(req.GetRole())
	if req.GetRole() == "" {
		return nil, util.ErrEmptyUserRole
	}
	if req.Role == user.Role {
		return user.ToProtoBuffer(), nil
	}

	user.Role = req.GetRole()
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	if err != nil {
		return nil, util.ErrUpdateUser
	}

	return user.ToProtoBuffer(), nil
}

// GetUserRole performs return the user role by id.
func (s *AuthService) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest) (*pb.GetUserRoleResponse, error) {
	if !bson.IsObjectIdHex(req.GetId()) {
		return nil, util.ErrInvalidUserId
	}

	user, err := s.usersRepository.GetById(req.GetId())
	if err != nil {
		return nil, util.ErrNotFoundUserId
	}

	return &pb.GetUserRoleResponse{Role: user.Role}, nil
}

// UpdateUser performs update the password.
func (s *AuthService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.User, error) {
	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}

	user, err := s.usersRepository.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}

	req.Password = strings.TrimSpace(req.GetPassword())
	if req.Password == "" {
		return nil, util.ErrEmptyPassword
	}

	req.NewPassword = strings.TrimSpace(req.GetNewPassword())
	if req.NewPassword == "" {
		return nil, util.ErrEmptyNewPassword
	}

	req.Password, err = security.EncryptPassword(req.GetPassword())
	if err != nil {
		return nil, util.ErrEncryptPassword
	}
	err = security.VerifyPassword(user.Password, req.GetPassword())
	if err != nil {
		return nil, util.ErrMismatchedPassword
	}
	req.NewPassword, err = security.EncryptPassword(req.GetNewPassword())
	if err != nil {
		return nil, util.ErrEncryptPassword
	}
	user.Password = req.GetNewPassword()
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	if err != nil {
		return nil, util.ErrUpdateUser
	}
	return user.ToProtoBuffer(), nil
}

// UpdateUser performs update the password.
func (s *AuthService) UpdateUserEmail(ctx context.Context, req *pb.UpdateUserEmailRequest) (*pb.User, error) {
	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}
	user, err := s.usersRepository.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}
	req.NewEmail = util.NormalizeEmail(req.GetNewEmail())
	if req.NewEmail == "" {
		return nil, util.ErrEmptyNewEmail
	}
	req.Password = strings.TrimSpace(req.GetPassword())
	if req.Password == "" {
		return nil, util.ErrEmptyPassword
	}
	req.Password, err = security.EncryptPassword(req.GetPassword())
	if err != nil {
		return nil, util.ErrEncryptPassword
	}
	err = security.VerifyPassword(user.Password, req.GetPassword())
	if err != nil {
		return nil, util.ErrMismatchedPassword
	}
	user.Email = req.GetNewEmail()
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	if err != nil {
		return nil, util.ErrUpdateUser
	}
	return user.ToProtoBuffer(), nil
}

// UpdateUser performs update the username.
func (s *AuthService) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest) (*pb.User, error) {
	normalizedEmail := util.NormalizeEmail(req.GetEmail())
	if normalizedEmail == "" {
		return nil, util.ErrEmptyEmail
	}
	user, err := s.usersRepository.GetByEmail(normalizedEmail)
	if err != nil {
		return nil, util.ErrNotFoundEmail
	}
	req.Password = strings.TrimSpace(req.GetPassword())
	if req.Password == "" {
		return nil, util.ErrEmptyPassword
	}
	req.Password, err = security.EncryptPassword(req.GetPassword())
	if err != nil {
		return nil, util.ErrEncryptPassword
	}
	err = security.VerifyPassword(user.Password, req.GetPassword())
	if err != nil {
		return nil, util.ErrMismatchedPassword
	}
	if strings.TrimSpace(req.GetName()) == "" {
		return nil, util.ErrEmptyName
	}
	user.Name = req.GetName()
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	if err != nil {
		return nil, util.ErrUpdateUser
	}
	return user.ToProtoBuffer(), nil
}

// ListUser list all users.
func (s *AuthService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error {
	users, err := s.usersRepository.GetAll()
	if err != nil {
		return util.ErrNotPerformedOperation
	}
	for _, user := range users {
		err := stream.Send(user.ToProtoBuffer())
		if err != nil {
			return util.ErrNotPerformedOperation
		}
	}
	return nil
}
