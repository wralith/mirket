package main

import (
	"context"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/wralith/mirket/server/internal/user"
	"github.com/wralith/mirket/server/pkg/valid"
)

func main() {
	os.Setenv("PG_CONN_STR", "postgresql://root:secret@postgres/mirket?sslmode=disable")
	config, err := initConfig()
	if err != nil {
		panic(err)
	}

	log.Println("auto migrate successful")
	autoMigrate(config.Postgres.ConnStr)

	pool, err := pgxpool.New(context.Background(), config.Postgres.ConnStr)
	if err != nil {
		panic(err)
	}

	repo := user.NewPostgresRepo(pool)
	svc := user.NewService(repo)
	cnt := user.NewHTTPController(svc)

	e := echo.New()
	e.Validator = &valid.Validator{Validator: validator.New()}

	g := e.Group("/users")
	user.InitHTTPEndpoints(g, cnt)

	e.Logger.Panic(e.Start(":" + config.App.Port))
}
