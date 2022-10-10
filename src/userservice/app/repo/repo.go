package repo

import "database/sql"

type Repo interface {
	Querier
}

type PgRepo struct {
	*Queries
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return &PgRepo{db: db, Queries: New(db)}
}
