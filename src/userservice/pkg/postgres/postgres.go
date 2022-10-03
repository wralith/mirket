package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(source string) *sql.DB {
	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	return db
}
