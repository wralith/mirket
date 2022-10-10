package grpcservice

import (
	"context"

	"github.com/rs/zerolog/log"
	pb "github.com/wralith/mirket/pb/user"
	"github.com/wralith/mirket/src/userservice/app/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	repo repo.Repo
}

func NewUserService(repo repo.Repo) *userService {
	return &userService{repo: repo}
}

func (s *userService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	user := req.GetUser()
	if user == nil {
		log.Error().Msg("unable to get user data from request body")
		return nil, status.Error(codes.Internal, "unable to get user data")
	}
	log.Debug().Msg(user.String())

	log.Debug().Msgf("AddUser request -> name:%s,email:%s,password:%s", user.Name, user.Email, user.Password)

	arg := &repo.AddUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	repoRes, err := s.repo.AddUser(ctx, *arg)
	if err != nil {
		log.Error().Err(err).Msg("unable to add user")
		return nil, err
	}

	res := mapToPb(&repoRes)
	return &pb.AddUserResponse{User: res}, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Debug().Msgf("GetUser request with id: %d", req.GetId())

	repoRes, err := s.repo.GetUser(ctx, req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("unable to get user")
		return nil, err
	}
	res := mapToPb(&repoRes)

	return &pb.GetUserResponse{User: res}, nil
}

func mapToPb(r *repo.User) *pb.User {
	return &pb.User{
		Id:       r.ID,
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Bio:      r.Bio,
	}
}
