package grpcServer

import (
	"context"

	"github.com/wralith/mirket/src/userservice/internal/user/usecase"
	"github.com/wralith/mirket/src/userservice/pb"
	"github.com/wralith/mirket/src/userservice/pkg/logger"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	uc     *usecase.UserUsecase
	logger *logger.Logger
}

func NewUserService(uc *usecase.UserUsecase, logger *logger.Logger) *userService {
	return &userService{uc: uc, logger: logger}
}

func (s *userService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	res, err := s.uc.AddUser(req.GetUser())
	if err != nil {
		s.logger.Errorln("error at add user grpc call")
		return nil, err
	}
	return &pb.AddUserResponse{User: res}, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := s.uc.GetUser(req.GetId())
	if err != nil {
		s.logger.Errorln("error at get user grpc call")
		return nil, err
	}
	return &pb.GetUserResponse{User: res}, nil
}
