package api

import (
	"github.com/rs/zerolog/log"
	"github.com/wralith/mirket/src/api-gateway/app/config"
	"github.com/wralith/mirket/src/api-gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUserServiceClient(c config.ServicesConfig) pb.UserServiceClient {
	conn, err := grpc.Dial(c.User, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("unable to dial user service")
	}

	client := pb.NewUserServiceClient(conn)
	return client
}
