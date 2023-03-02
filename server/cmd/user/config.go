package main

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Port string `env:"PORT,default=8080"`
}

type PostgresConfig struct {
	ConnStr string `env:"PG_CONN_STR,required"`
}

func initConfig() (config Config, err error) {
	err = envconfig.Process(context.Background(), &config)
	return
}
