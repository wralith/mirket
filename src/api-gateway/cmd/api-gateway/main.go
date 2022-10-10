package main

import (
	"github.com/wralith/mirket/src/api-gateway/app/api"
	"github.com/wralith/mirket/src/api-gateway/app/config"
	"github.com/wralith/mirket/src/api-gateway/app/handler"
	"github.com/wralith/mirket/src/api-gateway/app/logger"
	"github.com/wralith/mirket/src/api-gateway/app/router"
)

func main() {
	c := config.NewConfig()
	logger.InitLogger(&c.Logger)
	r := router.New()

	userSrv := api.CreateUserServiceClient(c.Services)
	userH := handler.NewUserHandler(userSrv)

	r.GET("/users/:id", userH.GetUser)

	panic(r.Start(":" + c.Server.Port))
}
