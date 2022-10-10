package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	pb "github.com/wralith/mirket/pb/user"
	"github.com/wralith/mirket/src/userservice/app/config"
	"github.com/wralith/mirket/src/userservice/app/grpcservice"
	"github.com/wralith/mirket/src/userservice/app/logger"
	"github.com/wralith/mirket/src/userservice/app/repo"
	"github.com/wralith/mirket/src/userservice/db/postgres"
	"google.golang.org/grpc"
)

func main() {
	c := config.NewConfig()
	logger.InitLogger(&c.Logger)

	db := postgres.Connect(c.Postgres)
	repo := repo.NewRepo(db)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Server.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start connection at port " + c.Server.Port)
	}

	service := grpcservice.NewUserService(repo)
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service)
	go server.Serve(conn)

	log.Info().Msg("Server started at port " + c.Server.Port)
	select {}
}
