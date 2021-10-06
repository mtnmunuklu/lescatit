package service

import (
	"Lescatit/authentication/models"
	"Lescatit/authentication/repository"
	"Lescatit/authentication/util"
	"Lescatit/pb"
	"Lescatit/security"
	"context"
	"log"
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
	normalizedEmail := util.NormalizeEmail(req.Email)
	found, err := s.usersRepository.GetByEmail(normalizedEmail)

	if err == mgo.ErrNotFound {
		user := new(models.User)
		user.Id = bson.NewObjectId()
		user.Name = strings.TrimSpace(req.Name)
		user.Email = normalizedEmail
		user.Password, err = security.EncryptPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Created = time.Now()
		user.Updated = time.Now()
		err := s.usersRepository.Save(user)
		if err != nil {
			return nil, err
		}
		return user.ToProtoBuffer(), nil
	}

	if found == nil {
		return nil, err
	}

	return nil, util.ErrEmailAlreadyExist
}

// SignIn performs the user login process.
func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	req.Email = util.NormalizeEmail(req.Email)
	user, err := s.usersRepository.GetByEmail(req.Email)
	if err != nil {
		log.Println("signin failed:", err.Error())
		return nil, util.ErrSignInFailed
	}

	err = security.VerifyPassword(user.Password, req.Password)
	if err != nil {
		log.Println("signin failed:", err.Error())
		return nil, util.ErrSignInFailed
	}

	token, err := security.NewToken(user.Id.Hex())
	if err != nil {
		log.Println("signin failed:", err.Error())
		return nil, util.ErrSignInFailed
	}

	return &pb.SignInResponse{User: user.ToProtoBuffer(), Token: token}, nil
}

// GetUser performs return the user by id.
func (s *AuthService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, util.ErrInvalidUserId
	}
	found, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return found.ToProtoBuffer(), nil
}

// ListUser list all users.
func (s *AuthService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error {
	users, err := s.usersRepository.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		err := stream.Send(user.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateUser performs update the user.
func (s *AuthService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, util.ErrInvalidUserId
	}
	user, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, util.ErrEmptyName
	}
	if req.Name == user.Name {
		return user.ToProtoBuffer(), nil
	}
	user.Name = req.Name
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	return user.ToProtoBuffer(), err
}

// DeleteUser performs delete the user.
func (s *AuthService) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, util.ErrInvalidUserId
	}
	err := s.usersRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Id: req.Id}, nil
}
