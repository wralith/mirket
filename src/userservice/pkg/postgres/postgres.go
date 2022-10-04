package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/wralith/mirket/src/userservice/config"
	"github.com/wralith/mirket/src/userservice/pkg/logger"
)

func Connect(c config.PostgresConfig) *sql.DB {
	log := logger.NewLogger()

	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DB)
	log.Info(dbSource)
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	return db
}
