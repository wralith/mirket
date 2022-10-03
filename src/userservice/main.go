package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wralith/mirket/src/userservice/internal/user/repo"
	"github.com/wralith/mirket/src/userservice/internal/user/usecase"
	"github.com/wralith/mirket/src/userservice/pb"
	"github.com/wralith/mirket/src/userservice/pkg/postgres"
	"google.golang.org/grpc"
)

var (
	log      *logrus.Logger
	port     = "3636"
	dbSource = "postgresql://root:password@localhost:5432/user?sslmode=disable"
)

func main() {
	log = logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	db := postgres.Connect(dbSource)
	repo := repo.NewRepo(db)
	uc := usecase.NewUserUsecase(repo)
	service := &userService{uc: uc}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Infof("grpc server starting at port: %s", port)

	run(service)
	select {}
}

func run(service *userService) {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service)

	go server.Serve(conn)
}

type userService struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUsecase
}

func (s *userService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	res, err := s.uc.AddUser(req.GetUser())
	if err != nil {
		log.Errorln("error at add user grpc call")
		return nil, err
	}
	return &pb.AddUserResponse{User: res}, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := s.uc.GetUser(req.GetId())
	if err != nil {
		log.Errorln("error at get user grpc call")
		return nil, err
	}
	return &pb.GetUserResponse{User: res}, nil
}
