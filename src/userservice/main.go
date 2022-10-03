package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wralith/mirket/src/userservice/pb"
	"google.golang.org/grpc"
)

var (
	log  *logrus.Logger
	port = "3636"
)

func main() {
	log = logrus.New()
	log.Formatter = &logrus.TextFormatter{}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Infof("grpc server starting at port: %s", port)

	run()
	select {}
}

func run() {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	service := &userService{}

	pb.RegisterUserServiceServer(server, service)

	go server.Serve(conn)
}

type userService struct {
	pb.UnimplementedUserServiceServer
}

func (s *userService) AddUser(context.Context, *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	return &pb.AddUserResponse{User: &pb.User{}}, nil
}

func (s *userService) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{User: &pb.User{Id: 1, Name: "Wra", Email: "wra@wra.com", Bio: "It is wralith!!"}}, nil
	// return &pb.GetUserResponse{User: &pb.User{}}, nil
}
