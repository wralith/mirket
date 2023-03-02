package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wralith/mirket/server/internal/user"
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

	// TODO: Move to test
	repo := user.NewPostgresRepo(pool)
	dummyUser := user.NewUser(user.NewUserOpts{
		Username:       "wra",
		Email:          "wralith",
		HashedPassword: []byte("secret"),
	})

	err = repo.Create(context.Background(), dummyUser)
	if err != nil {
		panic(err)
	}
	user, err := repo.GetByUsername(context.Background(), dummyUser.Username)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)
}
