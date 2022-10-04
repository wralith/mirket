package main

import (
	"fmt"
	"net"

	"github.com/wralith/mirket/src/userservice/config"
	"github.com/wralith/mirket/src/userservice/internal/grpcServer"
	"github.com/wralith/mirket/src/userservice/internal/user/repo"
	"github.com/wralith/mirket/src/userservice/internal/user/usecase"
	"github.com/wralith/mirket/src/userservice/pb"
	"github.com/wralith/mirket/src/userservice/pkg/logger"
	"github.com/wralith/mirket/src/userservice/pkg/postgres"
	"google.golang.org/grpc"
)

var log logger.Logger

type App struct {
	config *config.Config
}

func main() {
	log = *logger.NewLogger()
	cfg := config.NewConfig()
	a := App{config: cfg}

	db := postgres.Connect(a.config.Postgres)

	repo := repo.NewRepo(db)
	uc := usecase.NewUserUsecase(repo)

	log.Infof("grpc server starting at port: %s", a.config.Server.Port)

	a.run(uc)
	select {}
}

func (a *App) run(uc *usecase.UserUsecase) {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", a.config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	service := grpcServer.NewUserService(uc, &log)
	pb.RegisterUserServiceServer(server, service)

	go server.Serve(conn)
}
