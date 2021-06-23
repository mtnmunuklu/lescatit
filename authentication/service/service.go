package service

import (
	"CWS/authentication/models"
	"CWS/authentication/repository"
	"CWS/authentication/validators"
	"CWS/pb"
	"context"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type authService struct {
	usersRepository repository.UsersRepository
}

func NewAuthService(usersRepository repository.UsersRepository) pb.AuthServiceServer {
	return &authService{usersRepository: usersRepository}
}

func (s *authService) SignUp(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := validators.ValidateSignUp(req)
	if err != nil {
		return nil, err
	}

	found, err := s.usersRepository.GetByEmail(req.Email)
	if err == mgo.ErrNotFound {
		user := new(models.User)
		user.FromProtoBuffer(req)
		err := s.usersRepository.Save(user)
		if err != nil {
			return nil, err
		}
		return user.ToProtoBuffer(), nil
	}

	if found == nil {
		return nil, err
	}

	return nil, validators.ErrEmailAlreadyExist
}

func (s *authService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, validators.ErrInvalidUserId
	}
	found, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	return found.ToProtoBuffer(), nil
}

func (s *authService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error {
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

func (s *authService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, validators.ErrInvalidUserId
	}
	user, err := s.usersRepository.GetById(req.Id)
	if err != nil {
		return nil, err
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, validators.ErrEmptyName
	}
	if req.Name == user.Name {
		return user.ToProtoBuffer(), nil
	}
	user.Name = req.Name
	user.Updated = time.Now()
	err = s.usersRepository.Update(user)
	return user.ToProtoBuffer(), err
}

func (s *authService) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	if !bson.IsObjectIdHex(req.Id) {
		return nil, validators.ErrInvalidUserId
	}
	err := s.usersRepository.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Id: req.Id}, nil
}
